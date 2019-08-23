package main

import (
	"net/http"
	"fmt"
)

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("tajuden-token", "this is a secret key, keep it safe")
	res.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(res, "<h1>Any code you want in this func</h1>")
}

func main()  {
	var d hotdog
	http.ListenAndServe(":9080", d)
}