package main

import "fmt"

//* Interface como tipo genérico
//! cuidado para não virar gambiarra
func main() {
	generica(12)
	generica("Generica")
	generica(true)
	generica(18.95)
	generica(nil)
}

//Função com interface generica
func generica(interf interface{}) {
	fmt.Println(interf)
}
