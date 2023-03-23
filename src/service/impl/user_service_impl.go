package impl

import (
	"app/src/dao"
)

type UserServiceImpl struct {
}

func (UserServiceImpl) User(userId int) *dao.UserDao {
	user := &dao.UserDao{}
	user.SelectById(userId)
	return user
}
