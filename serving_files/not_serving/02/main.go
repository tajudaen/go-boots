package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.ListenAndServe(":9080", nil)
}

func dog(res http.ResponseWriter, req *http.Request)  {
	res.Header().Set("Content-Type", "text/html; charset=uft-8")

	io.WriteString(res, `
	<!--image doesn't serve -->
	<img src="/dummy.jpg">
	`)
}