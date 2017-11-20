package repository

import (
	"time"

	"bitbucket.org/DanielFrag/gestor-de-ponto/model"
	"gopkg.in/mgo.v2/bson"
)

//InsertDateRegister insert a date register
func InsertDateRegister(register model.DateRegister) error {
	s := getSession()
	defer s.Close()
	dateRegisterCollection := s.DB("dateRegister").C("DateRegisters")
	return dateRegisterCollection.Insert(register)
}

//GetUserRegistersByDate find documents by date
func GetUserRegistersByDate(userID bson.ObjectId, startDate time.Time, finishDate time.Time) ([]interface{}, error) {
	s := getSession()
	defer s.Close()
	dateRegisterCollection := s.DB("dateRegister").C("DateRegisters")
	var registers []interface{}
	err := dateRegisterCollection.
		Find(bson.M{
			"userId": userID,
			"timeStamp": bson.M{
				"$gte": startDate,
				"$lt":  finishDate,
			},
		}).
		All(&registers)
	return registers, err
}

//RemoveDateRegister find documents by date
func RemoveDateRegister(registerID bson.ObjectId) error {
	s := getSession()
	defer s.Close()
	dateRegisterCollection := s.DB("dateRegister").C("DateRegisters")
	return dateRegisterCollection.RemoveId(registerID)
}
