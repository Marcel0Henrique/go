package main

import "fmt"

//* Ponteiros
//? O ponteiro é uma referencia de memoria
func main() {
	// Variáveis são armazenadas na memoria e ganham um endereço na memoria
	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)
	// Mesmo as variáveis tendo o mesmo valor e sendo o valor de uma armazenado na outra
	// o endereço na memoria é diferente

	variavel1++
	fmt.Println(variavel1, variavel2)

	//* criando ponteiro
	var ponteiro *int // Ponteiro
	var ponteiro2 *int
	var variavel3 int

	variavel3 = 100
	ponteiro = &variavel3 // se coloca o "&" para passar a referencia da memoria da variavel para o ponteiro

	fmt.Println(ponteiro, ponteiro2, variavel3)
	fmt.Println("Endereço de memoria: ", ponteiro, " | Valor da variavel na memoria: ", *ponteiro)
	// Colocando "*" na frente do ponteiro mostra o seu valor
}
