package main

import "fmt"

//* Função com ponteiros
func main() {
	numero := 20
	numeroInvertido := InverterSinal(numero)
	fmt.Println(numeroInvertido)

	num := 40
	InverterSinalComPonteiro(&num)
	fmt.Println(num)
}

func InverterSinalComPonteiro(num *int) {
	*num = *num * -1
}

func InverterSinal(num int) int {
	return num * -1
}
