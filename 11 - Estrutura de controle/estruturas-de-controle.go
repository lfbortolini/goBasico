package main

import "fmt"

func main() {
	fmt.Println("Estrutura de controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual que 15")
	}

	if outronumero := numero; outronumero > 0 {
		fmt.Println("Numero é maior que zero")
	}
}
