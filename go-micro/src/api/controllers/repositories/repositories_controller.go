package repositories

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tajud99n/go-micro/src/api/domain/repositories"
	"github.com/tajud99n/go-micro/src/api/services"
	"github.com/tajud99n/go-micro/src/api/utils/errors"
)

func CreateRepo(c *gin.Context) {
	var request repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	result, err := services.RepositoryService.CreateRepo(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, result)
}
