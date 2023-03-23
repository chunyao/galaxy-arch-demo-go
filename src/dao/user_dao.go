package dao

import (
	"app/src/common/config/db"
	"app/src/model"
)

type UserDao struct {
	UserModel model.User
}

func (user *UserDao) SelectById(userId int) {
	db.RDBs["db2"].Db.First(&user.UserModel, userId)
}
