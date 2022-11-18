package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradoro string
	numero    uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Luiz"
	u.idade = 33

	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos Bobos", 0}

	usuario2 := usuario{"Luiz", 33, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Luiz"}
	fmt.Println(usuario3)
}
