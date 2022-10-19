package main

import (
	"fmt"
	"time"
)

// * Select
// ? basicamente é um switch para channel
func main() {
	canal1, canal2 := make(chan string), make(chan string)

	// esse canal recebe um valor a cada meio segundo
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal1"
		}
	}()

	// esse canal recebe um valor a cada dois segundo
	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal2"
		}
	}()

	// os valores são impressões ao mesmo tempo, porem o canal1 deveria ser mais rápido
	// mas está com a velocidade do canal2
	for {
		// usando o select se o canal1 já tiver recebido o valor vai ser impresso
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}
	}

}
