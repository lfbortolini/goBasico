package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Println("Dentro do metodo salvar")

	fmt.Printf("Salvando os dados do Usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiordeidade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuário 1", 20}
	fmt.Println(usuario1)
	usuario1.salvar()

	maiorDeIdade := usuario1.maiordeidade()
	fmt.Println(maiorDeIdade)

	usuario1.fazerAniversario()
}
