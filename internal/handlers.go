package internal

import (
	"encoding/json"
	"fmt"
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

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("%s, %s", user.Login, user.Password)

	w.WriteHeader(http.StatusOK)
}

func (rt *Router) PostUserSignup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func (rt *Router) PostUserInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
