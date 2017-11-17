package controller

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"github.com/gorilla/context"
	"bitbucket.org/DanielFrag/gestor-de-ponto/utils"
)

func RecoverFunc(w http.ResponseWriter, r *http.Request) {
	r.Body.Close()
	recoverError := recover()
	if recoverError != nil {
		http.Error(w, fmt.Sprint("Error: ", recoverError), http.StatusInternalServerError)
	}
}

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
					context.Set(r, "token", validatedToken)
					next.ServeHTTP(w, r)
					return
				}
				middlewareError = tokenError
			}
		}
		if middlewareError != nil {
			http.Error(w, "Error reading access token: " + middlewareError.Error(), http.StatusInternalServerError)
		} else {
			http.Error(w, "Error: no token provided", http.StatusInternalServerError)
		}
		return
	})
}