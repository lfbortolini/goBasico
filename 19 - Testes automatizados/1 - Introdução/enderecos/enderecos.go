package enderecos

import "strings"

// TipoDeEndereco verifica se um endereço tem um tipo valido
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	primeirapalavraDoEndereco := strings.Split(strings.ToLower(endereco), " ")[0]

	enderecoTemUmTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeirapalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return primeirapalavraDoEndereco
	}

	return "Tipo Inválido"
}
