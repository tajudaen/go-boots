package services

import (
	"github.com/tajud99n/go-micro/mvc/domain"
	"github.com/tajud99n/go-micro/mvc/utils"
)

type usersService struct {}

var (
	UsersService usersService
)

func (u *usersService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userId)
}