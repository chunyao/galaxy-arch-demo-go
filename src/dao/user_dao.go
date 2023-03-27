package dao

import (
	"app/src/common/config/db"
	"app/src/common/config/mongo"
	"app/src/model"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
)

type UserDao struct {
	UserModel model.User
}

func (user *UserDao) SelectById(userId int) {
	db.DBs["db2"].Db.First(&user.UserModel, userId)
}

func (user *UserDao) SelectByIdFromMongo(userId int) model.User {
	user.UserModel.Id = userId
	err := mongo.MGOs["wishproduct"].MG.Collection("user").FindOne(context.TODO(), &user.UserModel).Decode(&user.UserModel)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user.UserModel)
	return user.UserModel
}

func (user *UserDao) SaveToMongo(req model.User) {
	fmt.Printf("SaveToMongo %s", req)
	_, err := mongo.MGOs["wishproduct"].MG.Collection("user").InsertOne(context.TODO(), req)
	if err != nil {
		fmt.Printf("SaveToMongo error %s", err)
	}

}
