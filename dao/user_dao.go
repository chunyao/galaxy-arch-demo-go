package dao

import (
	"mabang-arch-demo-go/common/config/db"
	"mabang-arch-demo-go/model"
)

type UserDao struct {
	UserModel model.User
}

func (user *UserDao) SelectById(userId int) {
	db.RDBs["db2"].Db.First(&user.UserModel, userId)
}
