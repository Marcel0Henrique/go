package main

import "fmt"

//* Array e Slices
//? Array => lista de valores
//? Slice => é baseado no array porem é mais flexível

func main() {

	// Criando array
	// A quantidade de posições é obrigatória se não vira um "Slice"
	var array1 [5]int // array com 5 posições do tipo inteiro
	array2 := [5]int{}
	array3 := [...]int{1, 2, 3, 4, 5} // usando "..." o tamanho do array é definido pela quantidade de itens que foram passados no inicio

	// Adicionando item no Array
	array1[0] = 10
	array1[1] = 15

	fmt.Println(array1, array2, array3)

	// Criando slice
	// O tamanho de um slice é dinâmico porem continua a obrigatoriedade do tipo
	var slice1 []int
	slice2 := []int{}
	slice3 := array3[1:3]

	// Adicionando item no slice
	slice1 = append(slice1, 18)

	fmt.Println(slice1, slice2, slice3)

	array3[1] = 25

	fmt.Println(slice3)
}
