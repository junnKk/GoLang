package main

import (
   "encoding/json"
   "log"
   "net/http"
   "sort"
   "strconv"

   "github.com/gorilla/mux"
   "github.com/unrolled/render" // 웹 서버 응답 구현
   "github.com/urfave/negroni" // 웹 핸들러 제공 패키지 : 로그, panic복구, 파일 서버 자동 지원 기능
)

var rd *render.Render

type Todo struct { // 할일 정보담는 구조체
   ID        int    `json:"id,omitempty"` // JSON 포맷으로 변환 옵션, omitempty 생략 가능함을 나타냄
   Name      string `json:"name"`
   Completed bool   `json:"completed,omitempty"`
}

var todoMap map[int]Todo
var lastID int = 0

func MakeWebHandler() http.Handler {
   rd = render.New()
   todoMap = make(map[int]Todo)
   mux := mux.NewRouter()
   mux.Handle("/", http.FileServer(http.Dir("public")))
   mux.HandleFunc("/todos", GetTodoListHandler).Methods("GET")
   mux.HandleFunc("/todos", PostTodoHandler).Methods("POST")
   mux.HandleFunc("/todos/{id:[0-9]+}", RemoveTodoHandler).Methods("DELETE")
   mux.HandleFunc("/todos/{id:[0-9]+}", UpdateTodoHandler).Methods("PUT")
   return mux
}

type Todos []Todo // ID로 정렬하는 인터페이스

func (t Todos) Len() int {
   return len(t)
}

func (t Todos) Swap(i, j int) {
   t[i], t[j] = t[j], t[i]
}

func (t Todos) Less(i, j int) bool {
   return t[i].ID > t[j].ID
}

// 할일 목록
func GetTodoListHandler(w http.ResponseWriter, r *http.Request) {
   list := make(Todos, 0)
   for _, todo := range todoMap {
      list = append(list, todo)
   }
   sort.Sort(list) //ID로 정렬하여 전체 목록 반환
   rd.JSON(w, http.StatusOK, list)
}

// 추가
func PostTodoHandler(w http.ResponseWriter, r *http.Request) {
   var todo Todo
   err := json.NewDecoder(r.Body).Decode(&todo)
   //JSON 데이터를 Todo 객체로 변환 (디코딩)
   if err != nil {
      log.Fatal(err)
      w.WriteHeader(http.StatusBadRequest)
      return
   }
   lastID++
   todo.ID = lastID
   todoMap[lastID] = todo // id 발급해서 맵에 넣어준다.
   rd.JSON(w, http.StatusCreated, todo) // 다시 HTTP 응답을 통해 클라에게 전달
}

type Success struct {
   Success bool `json:"success"`
}

// 삭제
func RemoveTodoHandler(w http.ResponseWriter, r *http.Request) {
   vars := mux.Vars(r) 
   id, _ := strconv.Atoi(vars["id"])
   if _, ok := todoMap[id]; ok {
      delete(todoMap, id)
      rd.JSON(w, http.StatusOK, Success{true}) // 성공여부 전달
   } else {
      rd.JSON(w, http.StatusNotFound, Success{false}) // 실패여부 전달
   }
}

// 수정
func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
   var newTodo Todo
   err := json.NewDecoder(r.Body).Decode(&newTodo)
   //JSON 데이터를 Todo 객체로 변환 (디코딩)
   if err != nil {
      log.Fatal(err)
      w.WriteHeader(http.StatusBadRequest)
      return
   }

   vars := mux.Vars(r)
   id, _ := strconv.Atoi(vars["id"])
   if todo, ok := todoMap[id]; ok {
      todo.Name = newTodo.Name
      todo.Completed = newTodo.Completed
      rd.JSON(w, http.StatusOK, Success{true}) // 성공여부 전달
   } else {
      rd.JSON(w, http.StatusBadRequest, Success{false}) // 실패여부 전달
   }
}

func main() {
   m := MakeWebHandler()
   n := negroni.Classic() // negroni 핸들러를 통해 요청이 올때마다 로그를 찍어줌
   n.UseHandler(m)

   log.Println("Started App")
   err := http.ListenAndServe(":3000", n)
   if err != nil {
      panic(err)
   }
}