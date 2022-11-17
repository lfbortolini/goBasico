package main

import "fmt"

func main() {
	var variavelString = "Variavel String"
	variavelStringInferida := "Variavel Inferidas"

	fmt.Println(variavelString)
	fmt.Println(variavelStringInferida)

	var (
		variavel3 string = "Variavel 3"
		variavel4 string = "Variavel 4"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variavel 5", "Variavel 6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
