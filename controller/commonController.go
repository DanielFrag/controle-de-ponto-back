package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"

	"bitbucket.org/DanielFrag/gestor-de-ponto/business"
	"bitbucket.org/DanielFrag/gestor-de-ponto/model"
	"bitbucket.org/DanielFrag/gestor-de-ponto/repository"
	"bitbucket.org/DanielFrag/gestor-de-ponto/utils"
	"github.com/gorilla/context"
)

//CorsSetup middleware to allow cross domain origin
func CorsSetup(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			if r.Method == "OPTIONS" {
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

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
				if tokenError == nil {
					validatedToken, _ := utils.EncodeToken(payload)
					context.Set(r, "tokenPayload", payload)
					context.Set(r, "token", validatedToken)
					next(w, r)
					return
				}
				middlewareError = tokenError
			}
		}
		if middlewareError != nil {
			http.Error(w, "Error reading access token: "+middlewareError.Error(), http.StatusForbidden)
		} else {
			http.Error(w, "Error: no token provided", http.StatusBadRequest)
		}
		return
	})
}

//UserMiddleware middleware to check the user session
func UserMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenPayload := context.Get(r, "tokenPayload")
		if tokenMap, ok := tokenPayload.(map[string]interface{}); ok {
			userID, idOk := tokenMap["ID"].(string)
			session, sessionOk := tokenMap["Session"].(string)
			if idOk && sessionOk {
				user, err := repository.GetUserByID(bson.ObjectIdHex(userID))
				if err == nil && user.Session == session {
					context.Set(r, "user", user)
					next(w, r)
					return
				}
			}
		}
		http.Error(w, "Error reading user session", http.StatusBadRequest)
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
	tkPayload, loginError := business.Authenticate(body)
	if loginError != nil {
		http.Error(w, "Error to authenticate "+loginError.Error(), http.StatusForbidden)
		return
	}
	userWithSession, userSessionError := business.GenerateUserSession(tkPayload)
	if userSessionError != nil {
		http.Error(w, "Error to generate session "+userSessionError.Error(), http.StatusInternalServerError)
		return
	}
	tokenString, tokenAssignError := utils.EncodeToken(userWithSession)
	if tokenAssignError != nil {
		http.Error(w, "Error to assign token "+tokenAssignError.Error(), http.StatusInternalServerError)
		return
	}
	m := map[string]interface{}{
		"token": tokenString,
		"login": tkPayload.Login,
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(utils.FormatJSON(m))
	return
}

//UserLogout perform user logout
func UserLogout(w http.ResponseWriter, r *http.Request) {
	defer RecoverFunc(w, r)
	user := context.Get(r, "user").(model.User)
	repository.UpdateUserSession(user.ID, utils.GenerateRandomAlphaNumericString(15))
	return
}

//RenewToken just send a valid token
func RenewToken(w http.ResponseWriter, r *http.Request) {
	str := context.Get(r, "token")
	m := map[string]interface{}{
		"token": str,
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(utils.FormatJSON(m))
	return
}
