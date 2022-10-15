package main

import (
	"fmt"
)

// * Estrutura de repetição
// ? Dentro do GO existe apenas uma estrutura de repetição o laço FOR
// ? Porem tem varias formas de implementar e receber resultados diferentes
func main() {

	// For com index
	i := 0
	for i < 10 {
		i++
		fmt.Println(i)
		//time.Sleep(time.Second)
	}

	fmt.Println("-----------")

	// For com index²
	for j := 11; j <= 20; j++ {
		fmt.Println(j)
		//time.Sleep(time.Second)
	}

	nomes := [3]string{
		"Marcelo",
		"Clara",
		"Karol",
	}

	// For com Range
	for indice, nome := range nomes {
		fmt.Println(indice, " - ", nome)
	}

	//! Loop infinito
	// for {}
}
