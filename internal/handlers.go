package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (rt *Router) PostUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func (rt *Router) GetUser(w http.ResponseWriter, r *http.Request) {
	//key := r.Header.Get("X-Token")
	//
	//if key != rt.cfg.Token {
	//	w.WriteHeader(http.StatusUnauthorized)
	//	return
	//}

	var req map[string]string
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for key, value := range req {
		log.Printf("%s: %s", key, value)
	}

	w.WriteHeader(http.StatusOK)
	//fmt.Fprintf(w, "Hello World!")
}

func (rt *Router) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
