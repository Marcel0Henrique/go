package router

import (
	"devbook/src/router/routers"

	"github.com/gorilla/mux"
)

// * Retornar as rotas configuradas
func InitRouter() *mux.Router {
	r := mux.NewRouter()
	return routers.Config(r)
}
