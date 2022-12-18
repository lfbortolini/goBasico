package enderecos

import "testing"

type cenarioTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	enderecoParaTeste := "Avenida Paulista"
	tipoEnderecoEsperado := "Avenida"
	tipoEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if tipoEnderecoRecebido != tipoEnderecoEsperado {
		t.Errorf("O tipo recebido é diferente do esperado. %s é diferente de %s", tipoEnderecoEsperado, tipoEnderecoRecebido)
	}
}

func TestTipoDeEnderecoVariado(t *testing.T) {
	cenarioTeste := []cenarioTeste{
		{"Rua ABC", "Rua"},
		{"Rodovia das Nações", "Rodovia"},
		{"Praça das Bandeiras", "Tipo Inválido"},
		{"RUA DAS MULHERES", "Rua"},
	}

	for _, cenario := range cenarioTeste {
		tipoEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)

		if tipoEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado. %s é diferente de %s", cenario.retornoEsperado, tipoEnderecoRecebido)
		}
	}
}
