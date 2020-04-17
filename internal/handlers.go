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
		log.Printf("ERR\t%v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("REQ\tLOGIN\t{%s, %s}", user.Login, user.Password)

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

	log.Printf("REQ\tSIGNUP\t{%s, %s}", user.Login, user.Email)

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
		log.Printf("ERR\t%v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("REQ \tINFO\t{%s, %s}", user.Login, user.Password)

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
	type ResUser struct {
		AccountInfo UserHttp `json:"account_info"`
	}

	js, err := json.Marshal(ResUser{AccountInfo: userFromDB})
	if err != nil {
		log.Printf("ERR\t%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if bytes, err := w.Write(js); err != nil {
		log.Printf("ERR\t%v", err)
		http.Error(w, "ERR\tcannot write json to response: "+err.Error(), http.StatusInternalServerError)
	} else {
		log.Printf("RESP\tINFO\twritten %d bytes in response", bytes)
	}
}
