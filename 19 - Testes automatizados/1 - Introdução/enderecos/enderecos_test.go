package enderecos

import "testing"

func TestTipoDeEndereco(t *testing.T) {
	enderecoParaTeste := "Avenida Paulista"
	tipoEnderecoEsperado := "Rua"
	tipoEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if tipoEnderecoRecebido != tipoEnderecoEsperado {
		t.Errorf("O tipo recebido é diferente do esperado. %s é diferente de %s", tipoEnderecoEsperado, tipoEnderecoRecebido)
	}
}
