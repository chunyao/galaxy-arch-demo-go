package service

import "mabang-arch-demo-go/dao"

type UserService interface {
	User(userId int) *dao.User
}
