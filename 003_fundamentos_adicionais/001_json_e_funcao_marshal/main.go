package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// para a conversão de um struc para json
// os campos são mapeados já para facilitar a conversão
type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

// * Json
func main() {
	//? Converte um struc ou um map para um json
	//json.Marshal()

	c := cachorro{Nome: "Rex", Raca: "Pitbull", Idade: 6}
	fmt.Println("Struct: ", c)

	if cachorroEmJSON, err := json.Marshal(c); err != nil {
		log.Fatalln(err)
	} else {
		//? o json está como uint8 em bytes
		fmt.Println("Json (em uint8):", cachorroEmJSON)

		//? Mostra a saida em json
		fmt.Println(bytes.NewBuffer(cachorroEmJSON))
	}

}
