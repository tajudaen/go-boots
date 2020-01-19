package services

import (
	"net/http"
	"testing"

	"github.com/tajud99n/go-micro/mvc/domain"
	"github.com/tajud99n/go-micro/mvc/utils"

	"github.com/stretchr/testify/assert"
)

var (
	//userDaoMock usersDaoMock

	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationError)
)

func init() {
	domain.UserDao = &usersDaoMock{}
}

type usersDaoMock struct{}

func (m *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userId)
}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			Message: "user 0 not found",
			StatusCode: http.StatusNotFound,
		}
	}
	user, err := UsersService.GetUser(0)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 not found", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			Id: 123, 
			FirstName: "John",
			LastName: "Doe", 
			Email: "john@doe.com",
		}, nil
	}
	user, err := UsersService.GetUser(0)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "john@doe.com", user.Email)
}
