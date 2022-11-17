package main

import "fmt"

func main() {
	// NUMEROS INTEIROS
	// int8 int16 int32 int64
	var numero int64 = -10000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//alias
	// int32 = rune
	var numero3 rune = 12456
	fmt.Println(numero3)

	// byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	// NUMEROS REAIS
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12345678.90
	fmt.Println(numeroReal2)

	numeroReal3 := 12345678.90
	fmt.Println(numeroReal3)

	// STRINGS
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)

	char := "A"
	fmt.Println(char)

	char1 := 'A'
	fmt.Println(char1)

	// BOOL
	var booleano1 bool
	fmt.Println(booleano1)

	var erro error
	fmt.Println(erro)
}
