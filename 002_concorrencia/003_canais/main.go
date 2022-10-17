package main

import (
	"fmt"
	"time"
)

// * Channel
// ? Permite uma melhor sincronização do que o Go Routines
// ? É um canal de comunicação, permite enviar e receber dados
func main() {
	// Criando canal
	// o string especifica o tipo de dados que pode ser enviado e recebido
	canal := make(chan string)

	go escrever("Olá Mundo!", canal)

	fmt.Println("Depois da função escrever")

	for {

		mensagem, canalAberto := <-canal // canal está esperando um valor
		//canal retorna o valor recebido e também se está aberto ou fechado
		if canalAberto {
			fmt.Println(mensagem)
		} else {
			break
		}
	}
	fmt.Println("Fim do Programa")
}

func escrever(texto string, canal chan string) {

	for i := 0; i < 5; i++ {
		//? Passando valor para um canal
		canal <- texto
		time.Sleep(time.Second)
	}

	//! Fechando o canal para evitar deadlock
	// Deadlock é quando não tem amis informações sendo enviadas para o canal
	close(canal)
}
