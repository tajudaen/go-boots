package domain

import (
	"fmt"
	"net/http"

	"github.com/tajud99n/go-micro/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "John", LastName: "Doe", Email: "john@doe.com"},
	}

	UserDao usersDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type usersDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct {}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	fmt.Println("we are accessing the database")
	if user, ok := users[userId]; ok {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not found",
	}

}
