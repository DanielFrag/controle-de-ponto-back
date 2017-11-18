package business

import (
	"bytes"
	"encoding/json"
	"errors"

	"bitbucket.org/DanielFrag/gestor-de-ponto/dto"
	"bitbucket.org/DanielFrag/gestor-de-ponto/model"
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

//authUser check in repository the user credentials
func authUser(client dto.Login) (dto.AuthUser, error) {
	if client.Login != "ze" || client.Pass != "ze" {
		return dto.AuthUser{}, errors.New("authentication error")
	}
	return dto.AuthUser{
		ID:      "1",
		Session: "1",
	}, nil
}

//FormatToken format the json with the user's token that will be send to him
func FormatToken(tokenString string) []byte {
	b := new(bytes.Buffer)
	m := map[string]string{
		"token": tokenString,
	}
	json.NewEncoder(b).Encode(m)
	return b.Bytes()
}

func ExtractNewuserFormBody(body []byte) (model.User, error) {
	var user model.User
	jsonError := json.Unmarshal(body, &user)
	return user, jsonError
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
