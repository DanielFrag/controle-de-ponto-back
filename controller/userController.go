package controller

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"bitbucket.org/DanielFrag/gestor-de-ponto/business"
	"bitbucket.org/DanielFrag/gestor-de-ponto/model"
	"bitbucket.org/DanielFrag/gestor-de-ponto/repository"
	"bitbucket.org/DanielFrag/gestor-de-ponto/utils"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

//RegisterDate register a custom timestamp
func RegisterDate(w http.ResponseWriter, r *http.Request) {
	body, bodyReadError := ioutil.ReadAll(r.Body)
	if bodyReadError != nil {
		http.Error(w, "Error reading body from requested URL "+bodyReadError.Error(), http.StatusInternalServerError)
		return
	}
	defer RecoverFunc(w, r)
	user := context.Get(r, "user").(model.User)
	register, registerError := business.ExtractCustomTimestampFromBody(body)
	if registerError != nil {
		http.Error(w, "Error to extract custom timestamp from body "+registerError.Error(), http.StatusBadRequest)
		return
	}
	dateRegister := model.DateRegister{
		UserID:      user.ID,
		Timestamp:   register.Timestamp,
		Description: register.Description,
	}
	registerTimestampError := repository.InsertDateRegister(dateRegister)
	if registerTimestampError != nil {
		http.Error(w, "Error to register timestamp "+registerTimestampError.Error(), http.StatusInternalServerError)
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

//GetUserRegistersByDate recover the registers by user, based on timestamp interval
func GetUserRegistersByDate(w http.ResponseWriter, r *http.Request) {
	user := context.Get(r, "user").(model.User)
	dates := mux.Vars(r)
	if dates["start"] == "" || dates["finish"] == "" {
		http.Error(w, "Error, invalid date interval", http.StatusBadRequest)
		return
	}
	start, startError := strconv.ParseInt(dates["start"], 10, 64)
	finish, finishError := strconv.ParseInt(dates["finish"], 10, 64)
	if startError != nil || finishError != nil || start > finish {
		http.Error(w, "Error, invalid date format", http.StatusBadRequest)
		return
	}
	start /= 1000
	finish /= 1000
	registers, registersError := repository.GetUserRegistersByDate(user.ID, time.Unix(start, 0), time.Unix(finish, 0))
	if registersError != nil {
		http.Error(w, "Error to recover the timestamp registers "+registersError.Error(), http.StatusInternalServerError)
		return
	}
	str := context.Get(r, "token")
	m := map[string]interface{}{
		"token":         str,
		"dateRegisters": registers,
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(utils.FormatJSON(m))
	return
}

//RemoveDateRegister remove a single date register by id
func RemoveDateRegister(w http.ResponseWriter, r *http.Request) {
	user := context.Get(r, "user").(model.User)
	vars := mux.Vars(r)
	if vars["id"] == "" {
		http.Error(w, "Error, invalid date id", http.StatusBadRequest)
		return
	}
	dateRegisterID := business.ExtractIDFromString(vars["id"])
	removeDateRegisterError := repository.RemoveDateRegister(dateRegisterID, user.ID)
	if removeDateRegisterError != nil {
		http.Error(w, "Error to remove the date register "+removeDateRegisterError.Error(), http.StatusInternalServerError)
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
