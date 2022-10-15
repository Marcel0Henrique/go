package main

import "fmt"

//* Função anonima
func main() {
	// Criando e executando função anonima
	func(texto string) {
		fmt.Println(texto)
	}("Função Anonima") // se coloca o parentese no final para executar a função
	// também pode passar o parâmetro dentro do parentese

	soma := func(n1, n2 int) int {
		return n1 + n2
	}(10, 15)

	fmt.Println(soma)

}
