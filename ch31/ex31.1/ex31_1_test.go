package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostTodo(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/todos", strings.NewReader(`{"name":"ddd", "date":""}`))

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusCreated, res.Code)
	data, _ := io.ReadAll(res.Body)
	log.Println(string(data))
}
