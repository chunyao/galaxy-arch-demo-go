package model

type UserDo struct {
	Id   int    `bson:"id"`
	Name string `bson:"name"`
}
