package dao

import "mabang-arch-demo-go/common/config/db"

type User struct {
	Id   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"size:50"`
}

func (user *User) SelectById(userId int) {
	db.RDBs["db2"].Db.First(&user, userId)
}
