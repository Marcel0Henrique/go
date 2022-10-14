package main

import "fmt"

//* Estruturas de controles
func main() {
	idade := 18

	// Usando if, else if e else
	if idade > 18 {
		fmt.Println("MAIOR DE IDADE")
	} else if idade == 18 {
		fmt.Println("POR POUCO")
	} else {
		fmt.Println("MENOR DE IDADE")
	}

	// Declarando variável usando "if init"
	if outraIdade := idade; outraIdade > 0 {
		//? A variável outraIdade é local e apenas acessível dentro da estrutura de decisão
		fmt.Println("Idade:", outraIdade)
	}

}
