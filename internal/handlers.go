package internal

import (
	"encoding/json"
	"log"
	"net/http"
)

func (rt *Router) PostUserLogin(w http.ResponseWriter, r *http.Request) {
	//key := r.Header.Get("X-Token")
	//
	//if key != rt.cfg.Token {
	//	w.WriteHeader(http.StatusUnauthorized)
	//	return
	//}

	var user UserHttp
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("%s, %s", user.Login, user.Password)

	if IsUserExists(&user) {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
	return
}

func (rt *Router) PostUserSignup(w http.ResponseWriter, r *http.Request) {
	var user UserHttp
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("%s, %s", user.Login, user.Email)

	if IsUserLoginOrEmailExists(&user) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	if InsertUser(&user) {
		w.WriteHeader(http.StatusCreated)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	return
}

func (rt *Router) PostUserInfo(w http.ResponseWriter, r *http.Request) {
	var user UserHttp
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//InsertUserFull(&user)

	if !IsUserExists(&user) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	userFromDB := GetUserInfo(&user)
	if userFromDB.Login == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userFromDB.Password = ""
	js, err := json.Marshal(userFromDB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
