package main

import (
	"fmt"
	"sync"
	"time"
)

// * Wait Group(Grupo de espera)
// sincroniza as funções no caso para que terminem de executar ao mesmo tempo
func main() {

	var waitGroup sync.WaitGroup

	// Quantidade de Go Routines que vai rodar ao mesmo tempo
	waitGroup.Add(4)

	// Go Routines sendo executadas com wait group
	go func() {
		escrever("Go Routine 1")
		// Informa para o wait group que terminou a GO routine
		waitGroup.Done()
	}()

	go func() {
		escrever("Go Routine 2")
		waitGroup.Done()
	}()

	go func() {
		escrever("Go Routine 3")
		waitGroup.Done()
	}()

	go func() {
		escrever("Go Routine 4")
		waitGroup.Done()
	}()

	// o wait group espera as Go routines terminarem antes de encerrar o programa
	waitGroup.Wait()

}

func escrever(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
