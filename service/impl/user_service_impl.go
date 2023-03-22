package impl

import "mabang-arch-demo-go/dao"

type UserServiceImpl struct {
}

func (UserServiceImpl) User(userId int) *dao.User {
	user := &dao.User{}
	user.SelectById(userId)
	return user
}
