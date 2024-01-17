// Fatalf interromperá o teste

package formas

import (
	"interfaces/formas"
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retangulo", func(t *testing.T) {
		ret := formas.Retangulo{Largura: 10, Altura: 20}
		areaEsperada := float64(200)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
			// t.Errorf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})

	t.Run("Circulo", func(t *testing.T) {
		circ := formas.Circulo{Raio: 20}

		areaEsperada := float64(math.Pi * 20 * 20)
		areaRecebida := circ.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})
}
