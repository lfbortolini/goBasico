package main

import "fmt"

func soma(numeros ...int) int {
	fmt.Println(numeros)

	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalsoma := soma(1, 2, 3, 4, 5, 6, 200, 102, 12, 13)
	fmt.Println(totalsoma)

	escrever("Olá mundo", 10, 2)
}
