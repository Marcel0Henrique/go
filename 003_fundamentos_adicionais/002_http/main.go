package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {

	//* Definindo porta do servidor
	port := 5000

	log.Printf("Servidor iniciou em http://localhost:%v", port)

	//* Declarando rotas
	http.HandleFunc("/usuarios", getUsuario)

	//* O servidor escuta as requisições
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))

}

func getUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Sem usuarios cadastrados"))
}
