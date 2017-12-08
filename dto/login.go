package dto

import "gopkg.in/mgo.v2/bson"

//Login struct data transfer object
type Login struct {
	Login string `json:"login"`
	Pass  string `json:"password"`
}

//AuthUser struct data transfer object
type AuthUser struct {
	ID      bson.ObjectId
	Session string
	Login   string
}
