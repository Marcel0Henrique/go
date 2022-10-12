package main

import "fmt"

//* Herança
//? Não existe herança no GO, essa é a forma mais proxima disso

type pessoa struct {
	nome  string
	idade int8
}

// forma mais proxima de  realiza ruma herança
type estudante struct {
	pessoa // "Herança"
	curso  string
}

func main() {
	p1 := pessoa{nome: "joão", idade: 15}
	e1 := estudante{p1, "T.I"}
	e2 := estudante{pessoa{"Marcos", 26}, "Português"}

	fmt.Println(p1, e1, e2)
	fmt.Println(e1.nome, e2.pessoa.nome)
}
