package main

import "fmt"

func main() {

	//* Aritméticos
	soma := 1 + 1
	sub := 2 - 1
	divisao := 4 / 2
	multi := 2 * 3
	restoDaDivisao := 10 % 2

	fmt.Println(soma, sub, divisao, multi, restoDaDivisao)

	//* Atribuição
	var atribuicao1 string = "Atribuição"
	atribuicao2 := "Atribuição 2"

	fmt.Println(atribuicao1, atribuicao2)

	//* Relacionais
	maior := 10 > 2
	maiorIgual := 5 >= 4
	menor := 2 < 10
	menorIgual := 3 <= 6
	igual := 10 == 10
	diferente := 1 != 2

	fmt.Println(maior, maiorIgual, menor, menorIgual, igual, diferente)

	//* Lógicos
	and := true && false
	or := true || false
	not := !true

	fmt.Println(and, or, not)

	//* Unários
	num := 5
	num++
	num += 10
	num *= 2
	num /= 2
	num--
	num -= 5
	fmt.Println(num)

}
