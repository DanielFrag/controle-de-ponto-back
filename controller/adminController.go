package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"bitbucket.org/DanielFrag/gestor-de-ponto/business"
	"bitbucket.org/DanielFrag/gestor-de-ponto/utils"
	"github.com/gorilla/context"
)

func AdminLogin(w http.ResponseWriter, r *http.Request) {
	body, bodyReadError := ioutil.ReadAll(r.Body)
	if bodyReadError != nil {
		http.Error(w, "Error reading body from requested URL " + bodyReadError.Error(), http.StatusInternalServerError)
		return
	}
	defer RecoverFunc(w, r)
	tkPayload, adminLoginError := business.Authenticate(body)
	if adminLoginError != nil {
		http.Error(w, "Error to authenticate " + adminLoginError.Error(), http.StatusInternalServerError)
		return
	}
	tokenString, tokenAssignError := utils.EncodeToken(tkPayload)
	if tokenAssignError != nil {
		http.Error(w, "Error to assign token " + tokenAssignError.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(business.FormatToken(tokenString))
	return
}

func InsertUser(w http.ResponseWriter, r *http.Request) {
	/*
	body, bodyReadError := ioutil.ReadAll(r.Body)
	if bodyReadError != nil {
		http.Error(w, "Error reading body from requested URL " + bodyReadError.Error(), http.StatusInternalServerError)
		return
	}
	*/
	str := context.Get(r, "token")
	defer RecoverFunc(w, r)
	fmt.Fprintln(w, str)
}