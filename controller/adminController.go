package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"bitbucket.org/DanielFrag/gestor-de-ponto/business"
	"bitbucket.org/DanielFrag/gestor-de-ponto/model"
	"bitbucket.org/DanielFrag/gestor-de-ponto/repository"
	"bitbucket.org/DanielFrag/gestor-de-ponto/utils"
	"github.com/gorilla/context"
)

//AdminUserValidatorMiddleware check the user level
func AdminUserValidatorMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if user, ok := context.Get(r, "user").(model.User); ok {
			if user.Level == 1 {
				next(w, r)
				return
			}
		}
		http.Error(w, "Error, unauthorized user ", http.StatusBadRequest)
		return
	})
}

//InsertUser register a new user on application
func InsertUser(w http.ResponseWriter, r *http.Request) {
	body, bodyReadError := ioutil.ReadAll(r.Body)
	if bodyReadError != nil {
		http.Error(w, "Error reading body from requested URL "+bodyReadError.Error(), http.StatusInternalServerError)
		return
	}
	defer RecoverFunc(w, r)
	var user model.User
	var userError error
	userError = json.Unmarshal(body, &user)
	if userError != nil {
		http.Error(w, "Error reading user from body "+userError.Error(), http.StatusInternalServerError)
		return
	}
	userError = business.ValidateNewUserParams(user)
	if userError != nil {
		http.Error(w, "Error to validate user fields "+userError.Error(), http.StatusBadRequest)
		return
	}
	if business.CheckNewUser(user) {
		user.Password = business.GetDefaultPassword()
		repositoryError := repository.InsertUser(user)
		if repositoryError != nil {
			http.Error(w, "Error reading user from body "+repositoryError.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Error, the current user already exists", http.StatusBadRequest)
		return
	}
	str := context.Get(r, "token")
	m := map[string]interface{}{
		"token": str,
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(utils.FormatJSON(m))
	return
}
