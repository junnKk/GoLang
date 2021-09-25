package task

import (
	"errors"
	"fmt"
)

// InMemoryAccessor is a simple in-memory database.
type InMemoryAccessor struct {
	tasks  map[ID]Task
	nextID int64
}

// NewInMemoryAccessor returnsa new MemoryDataAccess.
func NewInMemoryAccessor() Accessor {
	return &InMemoryAccessor{
		tasks:  map[ID]Task{},
		nextID: int64(1),
	}
}

// ErrTaskNoExist occurs when the task with the ID was not found.
var ErrTaskNoExist = errors.New("task does not exist")

// Get returns a task with a given ID.
func (m *InMemoryAccessor) Get(id ID) (Task, error) {
	t, exists := m.tasks[id]
	if !exists {
		return Task{}, ErrTaskNoExist
	}
	return t, nil
}

// Put updates a task with a given ID with t.
func (m *InMemoryAccessor) Put(id ID, t Task) error {
	if _, exists := m.tasks[id]; !exists {
		return ErrTaskNoExist
	}
	m.tasks[id] = t
	return nil
}

// Post adds a new task
func (m *InMemoryAccessor) Post(t Task) (ID, error) {
	id := ID(fmt.Sprint(m.nextID))
	m.nextID++
	m.tasks[id] = t
	return id, nil
}

// Delete removes the task with a given ID
func (m *InMemoryAccessor) Delete(id ID) error {
	if _, exists := m.tasks[id]; !exists {
		return ErrTaskNoExist
	}
	delete(m.tasks, id)
	return nil
}
