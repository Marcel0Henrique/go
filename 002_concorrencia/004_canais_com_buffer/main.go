package main

import "fmt"

//* Canais com Buffer
func main() {
	// Criando canal com buffer, o canal tem capacidade de 2
	canal := make(chan string, 2)

	canal <- "Olá Mundo!"
	canal <- "Usando Go"

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem, mensagem2)
}
