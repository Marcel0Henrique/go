package main

import "fmt"

//* Função variaticas
//? pode receber vários parâmetros por padrão
func main() {
	totalSoma := soma(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(totalSoma)
}

// função vai receber vários parâmetros int
func soma(numeros ...int) (total int) {
	total = 0
	for _, numero := range numeros {
		total += numero
	}
	return
}
