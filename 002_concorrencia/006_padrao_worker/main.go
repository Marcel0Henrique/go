package main

import "fmt"

//* Padrão Worker Pool
func main() {

	// Canal com buffer de 45 posições
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	// processos sendo executada ao mesmo tempo para agilizar
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}

}

// especificando o a função de cada canal
// tarefas apenas recebe dados
// resultados apenas envia dados
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonnaci(numero)
	}
}

func fibonnaci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonnaci(posicao-2) + fibonnaci(posicao-1)
}
