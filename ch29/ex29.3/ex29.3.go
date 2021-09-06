//ch29/ex29.3/ex29.3.go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux() // serveMux 인스턴스 생성
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello world") // 인스턴스에 핸들러 등록

	})
	mux.HandleFunc("/bar", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello bar") // 인스턴스에 핸들러 등록

	})

	http.ListenAndServe(":3000", mux) // mux 인스턴스 사용
}
