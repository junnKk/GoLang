package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/jaeyeom/gogo/task"
)

type status int

const (
	UNKNOWN status = iota
	TODO
	DONE
)

func (s status) String() string {
	switch s {
	case UNKNOWN:
		return "UNKNOWN"
	case TODO:
		return "TODO"
	case DONE:
		return "DONE"
	default:
		return ""
	}
}

// Deadline is a struct to hold the deadline time.
type Deadline struct {
	time.Time
}

// NewDeadline returns a newly created Deadline with time t.
func NewDeadline(t time.Time) *Deadline {
	return &Deadline{t}
}

// ID is a data type to identify a task
type ID int64

// DataAccess is an interface to access task
type DataAccess interface {
	Get(id ID) (task.Task, error) // 어떤 ID를 주고 이 ID에 해당하는 작업을 추가해달라는 것
	Put(id ID, t task.Task) error // 이 ID의 작업을 넘겨준 t의 내용을 바꿔달라는 것
	Post(t task.Task) (ID, error) // 새로운 작업을 추가해달라는 것
	Delete(id ID) error           // 넘겨준 ID의 작업을 삭제해달라는 것
}

// MemoryDataAccess is a simple in-memory database.
type MemoryDataAccess struct {
	tasks  map[ID]task.Task
	nextID ID
}

// NewMemoryAccess returns a new MemoryDataAccess.
func NewMemoryDataAccess() DataAccess {
	return &MemoryDataAccess{
		tasks:  map[ID]task.Task{},
		nextID: ID(1),
	}
}

// ErrTaskNoExist occurs when the task with the ID was not found.
var ErrTaskNoExist = errors.New("task does not exist")

// Get returns a task with a given ID.
func (m *MemoryDataAccess) Get(id ID) (task.Task, error) {
	t, exists := m.tasks[id]
	if !exists {
		return task.Task{}, ErrTaskNoExist
	}
	return t, nil
}

// Put updates a task with a given ID with t.
func (m *MemoryDataAccess) Put(id ID, t task.Task) error {
	if _, exists := m.tasks[id]; !exists {
		return ErrTaskNoExist
	}
	m.tasks[id] = t
	return nil
}

// Post adds a new task
func (m *MemoryDataAccess) Post(t task.Task) (ID, error) {
	id :=m.nextID
	m.nextID++
	m.tasks[id] = t
	return id, nil
}

// Delete removes the task with a given ID
func (m *MemoryDataAccess) Delete(id ID) error {
	if _, exists := m.tasks[id]; !exists {
		return ErrTaskNoExist
	}
	delete(m.tasks, id)
	return nil
}

// ResponseError is the error for the JSON Response.
type ResponseError struct {
	Err error
}

// MarshalJSON returns the JSON representation of the error.
func (s status) MarshalJSON() ([]byte, error) {
	str := s.String()
	if str == "" {
		return nil, errors.New("status.MarshalJSON: unknown value")
	}
	return []byte(fmt.Sprintf(`"%s"`, str)), nil
}

// UnmarshalJSON parses the JSON represenstation of the error.
func (err *ResponseError) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, v); err != nil {
		return err
	}
	if v == nil {
		err.Err = nil
		return nil
	}
	switch tv := v.(type) {
	case string:
		if tv == ErrTaskNoExist.Error() {
			err.Err = ErrTaskNoExist
			return nil
		}
		err.Err = errors.New(tv)
		return nil
	default:
		return errors.New("ResponseError unmarshal failed")
	}
}

// Response is a struct for the JSON response.
type Response struct {
	ID    ID            `json:"id,omitempty"`
	Task  task.Task     `json:"task"`
	Error ResponseError `json:"error"`
}

var m = task.NewInMemoryAccessor()

const pathPrefix = "/api/v1/task"

func apiHandler(w http.ResponseWriter, r *http.Request) {
	getID := func() (ID, error) {
		id := ID(strconv.Atoi(r.URL.Path[len(pathPrefix):]))
		if id == "" {
			return id, errors.New("apiHandler: ID is Empty")
		}
		return id, nil
	}
	getTasks := func() ([]task.Task, error) {
		var result []task.Task
		if err := r.ParseForm(); err != nil {
			return nil, err
		}
		encodedTasks, ok := r.PostForm["task"]
		if !ok {
			return nil, errors.New("task parameter expected")
		}
		for _, encodedTask := range encodedTasks {
			var t task.Task
			if err := json.Unmarshal([]byte(encodedTask), &t); err != nil {
				return nil, err
			}
			result = append(result, t)
		}
		return result, nil
	}
	switch r.Method {
	case "GET":
		id, err := getID()
		if err != nil {
			log.Println(err)
			return
		}
		t, err := m.Get(id)
		err = json.NewEncoder(w).Encode(Response{
			ID:    id,
			Task:  t,
			Error: ResponseError{err},
		})
		if err != nil {
			log.Println(err)
		}
	case "PUT":
		id, err := getID()
		if err != nil {
			log.Println(err)
			return
		}
		tasks, err := getTasks()
		if err != nil {
			log.Println(err)
			return
		}
		for _, t := range tasks {
			err = m.Put(id, t)
			err = json.NewEncoder(w).Encode(Response{
				ID:    id,
				Task:  t,
				Error: ResponseError{err},
			})
			if err != nil {
				log.Println(err)
				return
			}
		}
	case "POST":
		tasks, err := getTasks()
		if err != nil {
			log.Println(err)
			return
		}
		for _, t := range tasks {
			id, err := m.Post(t)
			err = json.NewEncoder(w).Encode(Response{
				ID:    id,
				Task:  t,
				Error: ResponseError{err},
			})
			if err != nil {
				log.Println(err)
				return
			}
		}
	case "DELETE":
		id, err := getID()
		if err != nil {
			log.Println(err)
			return
		}
		err = m.Delete(id)
		err = json.NewEncoder(w).Encode(Response{
			ID:    id,
			Error: ResponseError{err},
		})
		if err != nil {
			log.Println(err)
			return
		}
	}
}


func main() {
	http.HandleFunc(pathPrefix, apiHandler)
	log.Fatal(http.ListenAndServe(":8884", nil))
}
