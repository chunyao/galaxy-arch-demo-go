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
	var userDo model.UserDo
	userDo.Id = req.Id
	userDo.Name = req.Name
	user.SaveToMongo(userDo)
	return user
}
