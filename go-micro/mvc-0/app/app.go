package app

import (
	"net/http"

	"github.com/tajud99n/go-micro/mvc/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}