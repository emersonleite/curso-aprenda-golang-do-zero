package formas

import (
	"fmt"
	"math"
)

// Retangulo
type Retangulo struct {
	Altura  float64
	Largura float64
}

// Circulo
type Circulo struct{ Raio float64 }

// Forma
type Forma interface {
	Area() float64
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.Raio, 2.0)
}

func EscreverArea(f Forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.Area())
}
