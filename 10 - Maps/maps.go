package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Luiz",
		"sobrenome": "Bortolini",
	}

	fmt.Println(usuario)

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Luiz",
			"segundo":  "Fernando",
		},
		"curso": {
			"nome":   "Analise desenvolvimento sistemas",
			"campus": "Campos 1 - Univel",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Gêmeos",
	}

	fmt.Println(usuario2)
}
