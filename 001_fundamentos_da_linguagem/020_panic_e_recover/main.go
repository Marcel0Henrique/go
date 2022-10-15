package main

import "fmt"

//* Panic e Recover
//? Diferente do erro que você consegue prosseguir depois que trata o erro,
//? o panic mata a execução do programa
//? Recover antes de o panic matar a execução vai chamar as funções que tem defer
func main() {
	defer RecuperarExecucao()
	fmt.Println(ALunoEstaAProvado(6, 6))
	fmt.Println("Pós execução")
}

func ALunoEstaAProvado(n1, n2 float64) bool {
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A media é exatamente 6!")
}

func RecuperarExecucao() {
	if rec := recover(); rec != nil {
		fmt.Println("Tentativa de recuperar a execução")
	}
}
