package controller

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/context"
)

//InsertUser register a new user on application
func InsertUser(w http.ResponseWriter, r *http.Request) {
	body, bodyReadError := ioutil.ReadAll(r.Body)
	defer RecoverFunc(w, r)
	log.Println(string(body[:]))
	if bodyReadError != nil {
		http.Error(w, "Error reading body from requested URL "+bodyReadError.Error(), http.StatusInternalServerError)
		return
	}
	str := context.Get(r, "token")
	defer RecoverFunc(w, r)
	fmt.Fprintln(w, str)
}
