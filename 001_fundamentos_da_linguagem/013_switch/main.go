package main

import "fmt"

//* switch
func main() {
	dia := DiaDaSemana(30)
	fmt.Println(dia)
}

// Retorna o dia da semana em string
func DiaDaSemana(num int) string {
	switch num {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Número invalido"
	}
}

// Outra maneira de usar o switch
func DiaDaSemana2(num int) string {
	switch {
	case num == 1:
		return "Domingo"
	case num == 2:
		return "Segunda-Feira"
	case num == 3:
		return "Terça-Feira"
	case num == 4:
		return "Quarta-Feira"
	case num == 5:
		return "Quinta-Feira"
	case num == 6:
		return "Sexta-Feira"
	case num == 7:
		return "Sábado"
	default:
		return "Número invalido"
	}
}
