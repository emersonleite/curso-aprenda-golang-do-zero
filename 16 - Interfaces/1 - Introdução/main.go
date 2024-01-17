package main

import (
	"interfaces/formas"
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

}
