package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/context"
)

//InsertUser register a new user on application
func InsertUser(w http.ResponseWriter, r *http.Request) {
	body, bodyReadError := ioutil.ReadAll(r.Body)
	if bodyReadError != nil {
		http.Error(w, "Error reading body from requested URL "+bodyReadError.Error(), http.StatusInternalServerError)
		return
	}
	str := context.Get(r, "token")
	defer RecoverFunc(w, r)
	fmt.Fprintln(w, str)
}
