package main

import "fmt"

func main() {
	sum, sub := calculosMatematicos(20, 10)
	fmt.Println(sum, sub)
}

// soma e subtração são retornos nomeados
func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}
