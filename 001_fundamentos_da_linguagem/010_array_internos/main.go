package main

import "fmt"

//* Arrays Internos
func main() {
	slice := make([]int, 2, 5)

	fmt.Println(slice)
	fmt.Println(len(slice)) // Tamanho
	fmt.Println(cap(slice)) // Capacidade

}
