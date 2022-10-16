package main

import "fmt"

//* Métodos
//? São parecidos com funções mas não são a mesma coisa
//? método está sempre associado a algo, normalmente a um struct
func main() {
	p1 := pessoa{"Marcelo", 19}
	fmt.Println(p1)
	p1.mostrarDados()

	p1.aumentarIdade()
	p1.mostrarDados()
}

type pessoa struct {
	nome  string
	idade uint8
}

//* Criando método
//? (p pessoa) vincula ao struct esse método
//? o "u" é uma variável que é usada para referenciar outros campos do struct pessoa
func (p pessoa) mostrarDados() {
	fmt.Printf("\nNome: %s\n", p.nome)
	fmt.Printf("Idade: %d", p.idade)
}

//* Criado método para alterar valor dentro da pessoa
//? se usa um ponteiro e não precisa desfazer a referenciação
func (p *pessoa) aumentarIdade() {
	p.idade++
}
