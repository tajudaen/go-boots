package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"github.com/tajud99n/go-micro/mvc/utils"

	"github.com/tajud99n/go-micro/mvc/services"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)
	if err != nil {
		// Just return the Bad Request to the client
		apiErr := &utils.ApplicationError{
			Message:    "userId must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad request",
		}

		utils.RespondError(c, apiErr)
		return
	}

	user, apiErr := services.UsersService.GetUser(userId)
	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}

	// return user to client
	utils.Respond(c, http.StatusOK, user)
}
