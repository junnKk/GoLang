package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/glowingedge/discoveryGo/chapter6/taskman/task"
)

// ResponseError is the error for the JSON Response.
type ResponseError struct {
	Err  error
	Code int
}

// MarshalJSON returns the JSON representation of the error.
func (err ResponseError) MarshalJSON() ([]byte, error) {
	if err.Err == nil {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%v"`, err.Err)), nil
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
		if tv == task.ErrTaskNoExist.Error() {
			err.Err = task.ErrTaskNoExist
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
	ID    task.ID       `json:"id,omitempty"`
	Task  task.Task     `json:"task"`
	Error ResponseError `json:"error"`
}
