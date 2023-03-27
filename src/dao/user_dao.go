package dao

import (
	"app/src/common/config/db"
	"app/src/common/config/mongo"
	"app/src/model"
	"context"
	"fmt"
)

type UserDao struct {
	UserModel model.User
}

func (user *UserDao) SelectById(userId int) {
	db.DBs["db2"].Db.First(&user.UserModel, userId)
}

func (user *UserDao) SelectByIdFromMongo(userId int) {
	db.DBs["db2"].Db.First(&user.UserModel, userId)
}

func (user *UserDao) SaveToMongo(req model.User) {
	ctx := context.Background()
	fmt.Printf("SaveToMongo %s", req)
	_, err := mongo.MGOs["wishproduct"].MG.Collection("user").InsertOne(ctx, req)
	if err != nil {
		fmt.Printf("SaveToMongo error %s", err)
	}

}
