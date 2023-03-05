package auxiliar

import "fmt"

//? é boa pratica colocar um comentário em cima de uma função publica

//* função com letra Maiúscula é publica
// Escrever registra uma mensagem na tela
func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar")
}

//* função com letra minuscula é privada
// ler escreve uma mensagem
func ler() {
	fmt.Println("Lendo do pacote auxiliar")
}
