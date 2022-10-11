package main

import "fmt"

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	// Função aninhada
	f := func(text string) {
		fmt.Println(text)
	}

	f("Texto da função F")

	soma, subtracao := CalculosMatematicos(10, 20)
	fmt.Println("Soma: ", soma, "Subtração: ", subtracao)

	//? Usando apenas o resultado de um retorno

	soma2, _ := CalculosMatematicos(10, 20) // o "_" faz ser ignorado o segundo resultado

	fmt.Println(soma2)
}

// * função privada com retorno
func somar(n1 int, n2 int) int {
	return n1 + n2
}

//* Função publica com múltiplos retornos
func CalculosMatematicos(n1 int, n2 int) (int, int) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}
