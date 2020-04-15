package internal

import (
	"fmt"
	"net/http"
)

func (rt *Router) PostUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
