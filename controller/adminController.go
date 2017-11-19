package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"bitbucket.org/DanielFrag/gestor-de-ponto/business"
	"bitbucket.org/DanielFrag/gestor-de-ponto/repository"
	"github.com/gorilla/context"
)

//InsertUser register a new user on application
func InsertUser(w http.ResponseWriter, r *http.Request) {
	body, bodyReadError := ioutil.ReadAll(r.Body)
	if bodyReadError != nil {
		http.Error(w, "Error reading body from requested URL "+bodyReadError.Error(), http.StatusInternalServerError)
		return
	}
	defer RecoverFunc(w, r)
	user, userError := business.ExtractUserFormBody(body)
	if userError != nil {
		http.Error(w, "Error reading user from body "+userError.Error(), http.StatusInternalServerError)
		return
	}
	if business.CheckNewUser(user) {
		repositoryError := repository.InsertUser(user)
		if repositoryError != nil {
			http.Error(w, "Error reading user from body "+repositoryError.Error(), http.StatusInternalServerError)
			return
		}
	}
	str := context.Get(r, "token")
	fmt.Fprintln(w, str)
}
