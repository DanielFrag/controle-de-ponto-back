package business

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

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
	return CheckClientCredentials(user)
}

//CheckClientCredentials find the user on db and check its credentials
func CheckClientCredentials(client dto.Login) (dto.AuthUser, error) {
	user := repository.GetUserByLogin(client.Login)
	if user.Login == "" || user.Login != client.Login || user.Password != EncryptString(client.Pass) {
		return dto.AuthUser{}, errors.New("authentication error")
	}
	return dto.AuthUser{
		ID:    user.ID,
		Login: user.Login,
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

//ValidateNewUserParams check user fields
func ValidateNewUserParams(user model.User) error {
	var userError error
	if user.Level < 0 || user.Level > 1 {
		userError = errors.New("Level must be 0 or 1")
	} else if user.Login == "" {
		userError = errors.New("No credentials provided")
	}
	return userError
}

//CheckNewUser check if the login is unique
func CheckNewUser(user model.User) bool {
	newUser := repository.GetUserByLogin(user.Login)
	return newUser.Login != user.Login
}

//EncryptString encrypt string with sha256
func EncryptString(s string) string {
	h := sha256.New()
	return fmt.Sprintf("%s", hex.EncodeToString(h.Sum([]byte(s))))
}

//GetDefaultPassword return the default password
func GetDefaultPassword() string {
	return EncryptString("password")
}
