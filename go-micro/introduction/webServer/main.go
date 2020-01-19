package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, request *http.Request) {
		w.Write([]byte("hello world"))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}