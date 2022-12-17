package main

import "fmt"

func main() {
	func(texto string) {
		fmt.Println(texto)
	}("Passando parametro")

	func() {
		fmt.Println("Olá mundo")
	}()

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passado por parametro")

	fmt.Println(retorno)

}
