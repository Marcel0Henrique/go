package main

import "fmt"

func main() {
	//* declarando variável
	//? 1 - método
	var variavel1 string = "Variável 1"
	fmt.Println(variavel1)

	//? 2 - método
	variavel2 := "Variável 2"
	fmt.Println(variavel2)

	//? 3 - declarando mais de uma variavel
	var (
		variavel3 string = "variavel 3"
		variavel4 string = "variavel 4"
	)

	fmt.Println(variavel3, variavel4)

	//? 3 - declarando mais de uma variavel²
	variavel5, variavel6 := "variavel 5", "variavel 6"

	fmt.Println(variavel5, variavel6)

	const constante1 string = "constante 1"
	fmt.Println(constante1)

	// trocando valor das variaveis
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println("Variavel 5: "+variavel5, "| Variavel 6 :"+variavel6)

}
