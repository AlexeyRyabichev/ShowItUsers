package internal

import (
	"github.com/AlexeyRyabichev/ShowItGate"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Route struct {
	Name        string           `json:"name"`
	Method      string           `json:"method"`
	Pattern     string           `json:"pattern"`
	HandlerFunc http.HandlerFunc `json:"-"`
}

type Router struct {
	cfg    ShowItGate.NodeCfg
	routes []Route

	Router *mux.Router
}

func NewRouter(cfg ShowItGate.NodeCfg) *Router {
	router := Router{
		cfg: cfg,
	}
	router.routes = []Route{
		{
			Name:        "Post user",
			Method:      "POST",
			Pattern:     "/user",
			HandlerFunc: router.PostUser,
		},
	}
	router.initRouter()
	return &router
}

func (rt *Router) initRouter() {
	rt.Router = mux.NewRouter().StrictSlash(true)

	for _, route := range rt.routes {
		rt.addRoute(route)
	}

	rt.Router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf(
			"HANDLER NOT FOUND FOR REQUEST: %s %s",
			r.Method,
			r.RequestURI,
		)
		w.WriteHeader(http.StatusNotFound)
	})
}

func (rt *Router) addRoute(route Route) {
	var handler http.Handler
	handler = route.HandlerFunc
	handler = ShowItGate.Logger(handler, route.Name)

	rt.Router.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(handler)
}
