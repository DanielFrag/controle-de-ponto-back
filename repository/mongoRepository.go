package repository

import (
	"os"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

func GetSession() *mgo.Session {
	mongoUrl := os.Getenv("MONGO_URL")
	if mongoUrl == "" {
		mongoUrl = "mongodb://localhost:27017"
	}
	s, err := mgo.Dial(mongoUrl)
	if err != nil {
        panic(err)
	}
	return s
}

func FinishSession(s *mgo.Session) {
	s.Close()
}