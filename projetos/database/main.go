package main

import (
	"database/env"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	env.Setenv()
	//* Criando rotas para API
	route := mux.NewRouter()
	//? Definindo a porta
	port := ":5000"
	fmt.Printf("Server start in port %s", port)
	log.Fatalln(http.ListenAndServe(port, route))
}
