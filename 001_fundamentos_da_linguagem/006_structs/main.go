package main

import "fmt"

//* Structs
//? é o mais proximo de uma classe em GO já que não tem orientação a objeto no GO
// structs é uma coleção de campos

// Criando o struct, primeiro se coloca o type depois o nome e por ultimo o struct
type usuario struct {
	nome  string
	idade uint8
}

// struct dentro de struct
type jogador struct {
	nome  string
	nivel int8
	skill habilidade
}

type habilidade struct {
	primaria   string
	secundaria string
}

func main() {
	// 1° método declarando o struct
	var user usuario
	user.nome = "Marcelo Henrique"
	user.idade = 23
	fmt.Println(user)

	// 2° método declarando o struct
	user2 := usuario{"Maria", 24}
	user3 := usuario{nome: "Paulo", idade: 17}
	user4 := usuario{nome: "Milena"}
	user5 := usuario{idade: 15}
	fmt.Println(user2, user3, user4, user5)

	hab := habilidade{"Ataque", "Defesa"}

	player := jogador{nome: "Icaro", nivel: 16, skill: hab}

	fmt.Println(player)
}
