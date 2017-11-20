package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//User application model
type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Login    string        `bson:"login"`
	Password string        `bson:"password"`
	Session  string        `bson:"session"`
	Level    int           `bson:"level"`
}

//DateRegister application model
type DateRegister struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	UserID    bson.ObjectId `bson:"userId"`
	Timestamp time.Time     `bson:"timeStamp"`
}
