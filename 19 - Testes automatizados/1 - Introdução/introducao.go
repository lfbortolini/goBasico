package main

import (
	"fmt"
	"introducao/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Praça das rosas")
	fmt.Println(tipoEndereco)
}
