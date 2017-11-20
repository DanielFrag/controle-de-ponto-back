package controller

import (
	"io/ioutil"
	"net/http"
	"time"

	"bitbucket.org/DanielFrag/gestor-de-ponto/business"
	"bitbucket.org/DanielFrag/gestor-de-ponto/model"
	"bitbucket.org/DanielFrag/gestor-de-ponto/repository"
	"bitbucket.org/DanielFrag/gestor-de-ponto/utils"
	"github.com/gorilla/context"
)

//RegisterNowTimestamp register the now timestamp
func RegisterNowTimestamp(w http.ResponseWriter, r *http.Request) {
	defer RecoverFunc(w, r)
	user := context.Get(r, "user").(model.User)
	insertDateRegisterError := registerTimestamp(user, time.Now())
	if insertDateRegisterError != nil {
		http.Error(w, "Error to register timestamp "+insertDateRegisterError.Error(), http.StatusBadRequest)
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

//RegisterCustomTimestamp register a custom timestamp
func RegisterCustomTimestamp(w http.ResponseWriter, r *http.Request) {
	body, bodyReadError := ioutil.ReadAll(r.Body)
	if bodyReadError != nil {
		http.Error(w, "Error reading body from requested URL "+bodyReadError.Error(), http.StatusInternalServerError)
		return
	}
	defer RecoverFunc(w, r)
	user := context.Get(r, "user").(model.User)
	customTimestamp, customTimestampError := business.ExtractCustomTimestampFromBody(body)
	if customTimestampError != nil {
		http.Error(w, "Error to extract custom timestamp from body "+customTimestampError.Error(), http.StatusBadRequest)
		return
	}
	registerTimestampError := registerTimestamp(user, customTimestamp)
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

func registerTimestamp(user model.User, timestamp time.Time) error {
	dateRegister := model.DateRegister{
		UserID:    user.ID,
		Timestamp: timestamp,
	}
	return repository.InsertDateRegister(dateRegister)
}

//GetUserRegisterByDate recover the registers by user, based on timestamp interval
func GetUserRegisterByDate(w http.ResponseWriter, r *http.Request) {
	body, bodyReadError := ioutil.ReadAll(r.Body)
	if bodyReadError != nil {
		http.Error(w, "Error reading body from requested URL "+bodyReadError.Error(), http.StatusInternalServerError)
		return
	}
	defer RecoverFunc(w, r)
	user := context.Get(r, "user").(model.User)
	startTs, finishTs, extractTsError := business.ExtractTimestampIntervalFromBody(body)
	if extractTsError != nil {
		http.Error(w, "Error reading timestamp interval "+extractTsError.Error(), http.StatusInternalServerError)
		return
	}
	registers, registersError := repository.GetUserRegistersByDate(user.ID, startTs, finishTs)
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
