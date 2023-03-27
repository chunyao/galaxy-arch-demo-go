package dao

import (
	"app/src/common/config/db"
	"app/src/common/config/mongo"
	"app/src/model"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

type UserDao struct {
	UserModel   model.User
	UserDoModel model.UserDo
}

func (user *UserDao) SelectById(userId int) {
	db.DBs["db2"].Db.First(&user.UserModel, userId)
}

func (user *UserDao) SelectByIdFromMongo(userId int) model.UserDo {
	user.UserDoModel.Id = userId
	filer := bson.D{{"id", userId}}
	err := mongo.MGOs["wishproduct"].MG.Collection("user").FindOne(context.TODO(), filer).Decode(&user.UserDoModel)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user.UserDoModel)
	return user.UserDoModel
}

func (user *UserDao) SaveToMongo(req model.UserDo) {
	fmt.Printf("SaveToMongo %s", req)
	_, err := mongo.MGOs["wishproduct"].MG.Collection("user").InsertOne(context.TODO(), req)
	if err != nil {
		fmt.Printf("SaveToMongo error %s", err)
	}

}
