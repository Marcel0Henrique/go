package main

import (
	"devbook/src/router"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	//* Definindo a porta
	port := 5000
	fmt.Printf("Start api server in http://localhost:%v", port)

	//* Injetando as rotas
	r := router.InitRouter()

	//* Iniciando servidor
	log.Fatalln(http.ListenAndServe(":"+strconv.Itoa(port), r))
}
