package main

import (
	"fmt"
	"time"
)

// * Concorrência
// ? Concorrência é diferente de paralelismo
// ? na concorrência não necessariamente as tarefas são realizadas ao mesmo tempo
// ? se o PC tiver mais de um núcleo então pode ser que sejam executadas ao mesmo tempo
// ? um exemplo básico seria que as tarefas ficam revesando a sua execução quando se tem apenas 1 núcleo
func main() {
	//* GO Routine
	// usando go, a função é executada mas o programa não espera a função terminar para poder prosseguir
	go escrever("testando")
	escrever("Go Routines")
}

func escrever(texto string) {
	for i := 0; i < 10; i++ {
		fmt.Println(texto+" - ", i)
		time.Sleep(time.Second)
	}
}
