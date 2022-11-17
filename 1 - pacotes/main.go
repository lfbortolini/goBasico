package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")

	erro := checkmail.ValidateFormat("lfbortolini@gmail.com")
	fmt.Println(erro)
}
