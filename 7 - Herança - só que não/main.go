package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

// Coloca-se a struct diretamente, sem colocar um tipo
type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {

	p1 := pessoa{
		"Emerson",
		"Leite",
		20,
		172,
	}

	e1 := estudante{p1, "Engenharia", "UTPR"}

	fmt.Println(e1)

	// Acesso direto a campos
	fmt.Println(e1.nome)
	fmt.Println(e1.altura)

}
