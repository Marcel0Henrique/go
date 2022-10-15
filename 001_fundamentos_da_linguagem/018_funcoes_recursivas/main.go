package main

import "fmt"

// Função recursiva
func main() {

	posisao := uint(10)

	fmt.Println(Fibonacci(posisao))

}

func Fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return Fibonacci(posicao-2) + Fibonacci(posicao-1)

}
