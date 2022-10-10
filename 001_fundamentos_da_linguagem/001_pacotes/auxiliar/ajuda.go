package auxiliar

import "fmt"

// Ajuda escreve uma mensagem
func Ajuda() {
	fmt.Println("Escrevendo do pacote auxiliar (função ajuda)")

	//? Pode chamar uma função privada se for do mesmo pacote
	ler()
}
