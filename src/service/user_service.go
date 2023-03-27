package service

import (
	"app/src/dao"
	"app/src/model"
)

type UserService interface {
	User(userId int) *dao.UserDao
	UserMongo(userId int) *dao.UserDao
	SaveUserMongo(req model.User) *dao.UserDao
}
