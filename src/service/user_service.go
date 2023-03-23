package service

import (
	"app/src/dao"
)

type UserService interface {
	User(userId int) *dao.UserDao
}
