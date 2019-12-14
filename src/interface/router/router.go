package router

import (
	"map-friend/src/interface/handler"
	"map-friend/src/internal/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	GetMethod    = "GET"
	PostMethod   = "POST"
	PutMethod    = "PUT"
	DeleteMethod = "DELETE"
)

type Route struct {
	Path    string
	Methods []string
	Handler func(w http.ResponseWriter, r *http.Request)
}

func routes() []*Route {
	return []*Route{
		{
			Path:    "/",
			Methods: []string{GetMethod, PostMethod},
			Handler: handler.HelloHandler(),
		},
		{
			Path:    "/users/{id:[0-9]+}",
			Methods: []string{GetMethod},
			Handler: handler.GetUserHandler(),
		},
	}
}

func New() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.ReqIdMiddleware)
	for _, v := range routes() {
		r.HandleFunc(v.Path, v.Handler).Methods(v.Methods...)
	}
	return r
}
