package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Hello World")  // 웹 핸들러 등록
	})

	http.ListenAndServe(":3000", nil) // 웹서버 시작
}