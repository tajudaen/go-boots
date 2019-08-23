package main

import (
	"net/http"
	"io"
)

type hotdot int

func (m hotdog) ServeHTTP(res http.Responsewriter, req *http.Request) {
	switch req.URL.Path {
	case "/test":
		io.WriteString(res, "test test test")
	case "/practice":
		io.WriteString(res, "practice practice practice")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":9080", d)
}