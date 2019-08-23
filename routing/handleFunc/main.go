package main

import (
	"io"
	"net/http"
)

func d(res http.Responsewrite, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func c(res http.Responsewrite, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {

	mux.HandleFunc("/dog/", d)
	mux.HandleFunc("/cat", c)

	http.ListenAndServe(":9080", nil)
}