package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// * Router representa todas as rodas da API
type Router struct {
	Uri                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	AuthenticationRequired bool
}

func Config(r *mux.Router) *mux.Router {
	routes := userRouters

	for _, route := range routes {
		r.HandleFunc(route.Uri, route.Function).Methods(route.Method)
	}

	return r
}
