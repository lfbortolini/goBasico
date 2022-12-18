package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retangulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.Errorf("A área recebida está divergente. Esperada %f / Recebida %f", areaEsperada, areaRecebida)
		}
	})

	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{5}
		areaEsperada := float64(math.Pi * math.Pow(5, 2))
		areaRecebida := circ.Area()

		if areaEsperada != areaRecebida {
			t.Errorf("A área recebida está divergente. Esperada %f / Recebida %f", areaEsperada, areaRecebida)
		}
	})
}
