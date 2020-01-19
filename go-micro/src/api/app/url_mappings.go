package app

import (
	"github.com/tajud99n/go-micro/src/api/controllers/ping"
	"github.com/tajud99n/go-micro/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.POST("/repositories", repositories.CreateRepo)
}
