package app

import (
	"github.com/tajud99n/go-micro/mvc/controllers"
)
func mapUrls() {
	router.GET("/users/:userId", controllers.GetUser)
}