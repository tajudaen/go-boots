package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/tajud99n/go-micro/mvc/utils"

	"github.com/tajud99n/go-micro/mvc/services"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("userId"), 10, 64)
	if err != nil {
		// Just return the Bad Request to the client
		apiErr := &utils.ApplicationError{
			Message:    "userId must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad request",
		}
		jv, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.StatusCode)
		res.Write(jv)
		return
	}

	user, apiErr := services.UsersService.GetUser(userId)
	if apiErr != nil {
		// Handle the err and return to the client
		jv, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.StatusCode)
		res.Write(jv)
		return
	}

	// return user to client
	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)
}
