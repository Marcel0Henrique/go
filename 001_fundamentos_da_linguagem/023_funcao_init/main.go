package main

import "fmt"

// Variável global
var n int

//* Função Init
//? Função init é executada antes da função main
//? Função init pode ser uma por arquivo diferente da main
func init() {
	fmt.Println("Executando função init")
	n = 10
}

func main() {
	fmt.Println("Função main sendo executada")
	fmt.Println("Variável global: ", n)
}
