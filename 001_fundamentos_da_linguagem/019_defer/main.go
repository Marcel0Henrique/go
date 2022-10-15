package main

import "fmt"

//* Defer
//? Basicamente adia a execução de um bloco de código
func main() {
	defer funcao1()
	funcao2()
}

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}
