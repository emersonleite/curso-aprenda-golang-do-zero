package main

import (
	"fmt"
	"interfaces/formas"

	"github.com/emersonleite/golang_labs/fibonacci"
)

func main() {

	ret := formas.Retangulo{
		Altura:  10,
		Largura: 15,
	}

	circ := formas.Circulo{
		Raio: 10,
	}

	formas.EscreverArea(circ)
	formas.EscreverArea(ret)

	_, fibonacci := fibonacci.CalculateMany(10)

	fmt.Println(fibonacci)

}
