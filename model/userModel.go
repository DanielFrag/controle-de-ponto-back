package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id       bson.ObjectId `bson:"_id"`
	Login    string        `bson:"login"`
	Password string        `bson:"password"`
	Session  string        `bson:"session"`
	Level    int           `bson:"level"`
}

type DateRegister struct {
	Id        bson.ObjectId `bson:"_id"`
	UserId    bson.ObjectId `bson:"userId"`
	TimeStamp time.Time     `bson:"timeStamp"`
}
