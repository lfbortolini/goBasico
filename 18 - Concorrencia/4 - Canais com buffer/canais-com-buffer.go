package main

import (
	"fmt"
)

func main() {
	canal := make(chan string)
	canal <- "Olá mundo novo"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
