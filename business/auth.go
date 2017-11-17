package business

import (
	"bytes"
	"errors"
	"encoding/json"
	"bitbucket.org/DanielFrag/gestor-de-ponto/dto"
)

func Authenticate(body []byte) (dto.AuthClient, error) {
	var client dto.Login
	jsonError := json.Unmarshal(body, &client)
	if jsonError != nil {
		return dto.AuthClient{}, jsonError
	}
	return authClient(client)
}

func authClient(client dto.Login) (dto.AuthClient, error) {
	if client.Login != "ze" || client.Pass != "ze" {
		return dto.AuthClient{}, errors.New("authentication error")
	}
	return dto.AuthClient {
		ID: "1",
		Level: "1",
		Session: "1",
	}, nil
}

func FormatToken(tokenString string) []byte {
	b := new (bytes.Buffer)
	m := map[string]string {
		"token": tokenString,
	}
	json.NewEncoder(b).Encode(m)
	return b.Bytes()
}