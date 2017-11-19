package repository

import (
	"bitbucket.org/DanielFrag/gestor-de-ponto/model"
	"gopkg.in/mgo.v2/bson"
)

//InsertUser insert one user on DB
func InsertUser(user model.User) error {
	s := getSession()
	defer s.Close()
	usersCollection := s.DB("dateRegister").C("Users")
	return usersCollection.Insert(user)
}

//GetUserByLogin find one user by login
func GetUserByLogin(login string) model.User {
	s := getSession()
	defer s.Close()
	usersCollection := s.DB("dateRegister").C("Users")
	var user model.User
	usersCollection.
		Find(map[string]string{
			"login": login,
		}).
		One(&user)
	return user
}

//GetUserByID update the field session
func GetUserByID(userID bson.ObjectId) (model.User, error) {
	s := getSession()
	defer s.Close()
	usersCollection := s.DB("dateRegister").C("Users")
	var user model.User
	err := usersCollection.Find(bson.M{"_id": userID}).One(&user)
	return user, err
}

//UpdateUserSession update the field session
func UpdateUserSession(userID bson.ObjectId, session string) error {
	s := getSession()
	defer s.Close()
	usersCollection := s.DB("dateRegister").C("Users")
	return usersCollection.Update(bson.M{
		"_id": userID,
	}, bson.M{
		"$set": bson.M{
			"session": session,
		},
	})
}
