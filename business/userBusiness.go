package business

import (
	"encoding/json"
	"errors"
	"time"

	"bitbucket.org/DanielFrag/gestor-de-ponto/utils"

	"bitbucket.org/DanielFrag/gestor-de-ponto/dto"
	"bitbucket.org/DanielFrag/gestor-de-ponto/model"
	"bitbucket.org/DanielFrag/gestor-de-ponto/repository"
)

//Authenticate extract user from json and check his credentials
func Authenticate(body []byte) (dto.AuthUser, error) {
	var user dto.Login
	jsonError := json.Unmarshal(body, &user)
	if jsonError != nil {
		return dto.AuthUser{}, jsonError
	}
	return authUser(user)
}

func authUser(client dto.Login) (dto.AuthUser, error) {
	user := repository.GetUserByLogin(client.Login)
	if user.Login != client.Login {
		return dto.AuthUser{}, errors.New("authentication error")
	}
	return dto.AuthUser{
		ID: user.ID,
	}, nil
}

//GenerateUserSession save a random session
func GenerateUserSession(user dto.AuthUser) (dto.AuthUser, error) {
	var err error
	user.Session = utils.GenerateRandomAlphaNumericString(10)
	err = repository.UpdateUserSession(user.ID, user.Session)
	return user, err
}

//CheckUserSession verify the user session
func CheckUserSession(user dto.AuthUser) bool {
	userModel, userError := repository.GetUserByID(user.ID)
	if userError != nil || userModel.Session != user.Session {
		return false
	}
	return true
}

//ExtractUserFormBody return the User model extracted from requisition body
func ExtractUserFormBody(body []byte) (model.User, error) {
	var user model.User
	var err error
	err = json.Unmarshal(body, &user)
	if err == nil {
		err = isValidUserParams(user)
	}
	return user, err
}

func isValidUserParams(user model.User) error {
	var userError error
	if user.Level < 0 || user.Level > 1 {
		userError = errors.New("Level must be 0 or 1")
	} else if user.Login == "" || user.Password == "" {
		userError = errors.New("No credentials provided")
	}
	return userError
}

//CheckNewUser check if the login is unique
func CheckNewUser(user model.User) bool {
	newUser := repository.GetUserByLogin(user.Login)
	return newUser.Login != user.Login
}

//ExtractCustomTimestampFromBody extract the timestamp from json with key "customTimestamp"
func ExtractCustomTimestampFromBody(body []byte) (time.Time, error) {
	var jsonMap map[string]time.Time
	err := json.Unmarshal(body, &jsonMap)
	return jsonMap["customTimestamp"], err
}

//ExtractTimestampIntervalFromBody extract the timestamp interval from json
func ExtractTimestampIntervalFromBody(body []byte) (time.Time, time.Time, error) {
	var jsonMap map[string]time.Time
	err := json.Unmarshal(body, &jsonMap)
	return jsonMap["startTimestamp"], jsonMap["finishTimestamp"], err
}
