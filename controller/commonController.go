package controller

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"bitbucket.org/DanielFrag/gestor-de-ponto/business"
	"bitbucket.org/DanielFrag/gestor-de-ponto/utils"
	"github.com/gorilla/context"
)

//RecoverFunc is used by a controller's defer statement. It will close the request's body, check for general errors and format an error response
func RecoverFunc(w http.ResponseWriter, r *http.Request) {
	r.Body.Close()
	recoverError := recover()
	if recoverError != nil {
		http.Error(w, fmt.Sprint("Error: ", recoverError), http.StatusInternalServerError)
	}
}

//TokenCheckerMiddleware is used like a middleware. It perform a user's token validation, extracting the payload, and generate a new token for user
func TokenCheckerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var middlewareError error
		authHeader := r.Header.Get("authorization")
		if len(authHeader) > 0 {
			authHeaderData := strings.Split(authHeader, " ")
			if len(authHeaderData) > 1 {
				payload, tokenError := utils.TokenChecker(authHeaderData[1])
				log.Println(payload)
				if tokenError == nil {
					validatedToken, _ := utils.EncodeToken(payload)
					context.Set(r, "tokenPayload", payload)
					context.Set(r, "token", validatedToken)
					next.ServeHTTP(w, r)
					return
				}
				middlewareError = tokenError
			}
		}
		if middlewareError != nil {
			http.Error(w, "Error reading access token: "+middlewareError.Error(), http.StatusInternalServerError)
		} else {
			http.Error(w, "Error: no token provided", http.StatusInternalServerError)
		}
		return
	})
}

//UserLogin athenticate the user
func UserLogin(w http.ResponseWriter, r *http.Request) {
	body, bodyReadError := ioutil.ReadAll(r.Body)
	if bodyReadError != nil {
		http.Error(w, "Error reading body from requested URL "+bodyReadError.Error(), http.StatusInternalServerError)
		return
	}
	defer RecoverFunc(w, r)
	tkPayload, adminLoginError := business.Authenticate(body)
	if adminLoginError != nil {
		http.Error(w, "Error to authenticate "+adminLoginError.Error(), http.StatusInternalServerError)
		return
	}
	tokenString, tokenAssignError := utils.EncodeToken(tkPayload)
	if tokenAssignError != nil {
		http.Error(w, "Error to assign token "+tokenAssignError.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(business.FormatToken(tokenString))
	return
}
