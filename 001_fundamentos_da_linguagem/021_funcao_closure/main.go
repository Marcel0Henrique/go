package main

import "fmt"

//* Função closure
//? Basicamente são funções que referenciam variáveis que estão
//? fora do seu corpo
func main() {
	texto := "Dento da função main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}

func closure() (funcao func()) {
	texto := "Dentro da função closure"
	funcao = func() {
		fmt.Println(texto)
	}
	return
}
