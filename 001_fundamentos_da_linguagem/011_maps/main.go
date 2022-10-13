package main

import "fmt"

//* Maps
func main() {

	// criando map
	// dentro do "[]" é o tipo da chave, fora é o tipo do valor
	usuarios := map[string]string{
		"nome":      "pedro",
		"sobrenome": "silva",
	}

	delete(usuarios, "sobrenome")
	usuarios["signo"] = "Virgem"

	fmt.Println(usuarios, usuarios["nome"])
}
