package impl

import (
	"app/src/dao"
	"app/src/model"
)

type UserServiceImpl struct {
}

func (UserServiceImpl) User(userId int) *dao.UserDao {
	user := &dao.UserDao{}
	user.SelectById(userId)
	return user
}

func (UserServiceImpl) UserMongo(userId int) *dao.UserDao {
	user := &dao.UserDao{}
	user.SelectByIdFromMongo(userId)
	return user
}

func (UserServiceImpl) SaveUserMongo(req model.User) *dao.UserDao {
	user := &dao.UserDao{}
	user.SaveToMongo(req)
	return user
}
