package main

import (
	"errors"
	"fmt"
)

func main() {
	//* Tipos básicos

	//* Inteiros(positivo e negativo)
	var (
		//? apenas muda a quantidade de bytes que cada inteiro suporta
		inteiro   int   = 00 // não precisa especificar e utiliza a arquitetura do processador para especificar o tipo de int
		inteiro8  int8  = 8
		inteiro16 int16 = 16
		inteiro32 int32 = 32
		inteiro64 int64 = 64

		inteiroSemSinal uint = 2 // não permite numero negativo

		//* Alias
		//? INT32 = RUNE
		inteiro32Alias rune = 12345

		//? UINT8 = BYTE
		inteiroSemSinal8Alias byte = 128
	)
	fmt.Println("Inteiros: ", inteiro, inteiro8, inteiro16, inteiro32, inteiro64)
	fmt.Println("Inteiro sem sinal: ", inteiroSemSinal)
	fmt.Println("Inteiro Alias: ", inteiro32Alias)
	fmt.Println("Inteiro Sem Sinal Alias: ", inteiroSemSinal8Alias)

	//* Float(Número reais)
	var (
		flutuante32 float32 = 32.648547
		flutuante64 float64 = 64.784525
	)

	fmt.Println("Flutuantes: ", flutuante32, flutuante64)

	//* String(Texto)
	var texto string = "Olá Mundo!"
	fmt.Println("Texto: " + texto)

	//* Valor zero(Valor inicial)
	var textoZero string
	var inteiroZero int
	fmt.Println(textoZero)
	fmt.Println(inteiroZero)

	//* Bool(Boleano)
	var verdadeiro bool = true
	var falso bool = false

	fmt.Println("Boleano: ", verdadeiro, falso)

	//* Error
	var erro error = errors.New("Novo Erro Criado")
	var erroZero error
	fmt.Println("Erro: ", erro, erroZero)
}
