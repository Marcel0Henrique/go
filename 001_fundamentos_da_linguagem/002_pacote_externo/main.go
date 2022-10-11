package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	erro := checkmail.ValidateFormat("teste@hotmail.com")
	fmt.Println(erro)
}
