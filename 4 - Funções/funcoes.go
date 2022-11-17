package main

import "fmt"

func main() {
	fmt.Println(somar(2, 5))

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto passado para função f")
	fmt.Println(resultado)

	resultadosSoma, resultadoSubtracao := calculosMatematicos(5, 10)
	fmt.Println(resultadosSoma, resultadoSubtracao)

	_, resultadoSubtracao1 := calculosMatematicos(5, 10)
	fmt.Println(resultadoSubtracao1)
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func somar(n1, n2 int8) int8 {
	return n1 + n2
}
