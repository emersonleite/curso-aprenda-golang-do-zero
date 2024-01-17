package main

import (
	"fmt"
)

func main() {
	//  For com condição somente
	/* i := 0
	for i < 10 {
		i++
		fmt.Println("incrementando i", i)
		time.Sleep(time.Second)
	} */

	// For com inicialização, condição  e incremento
	/* for j := 0; j < 10; j++ {
		fmt.Println("incrementando j", j)
		time.Sleep(time.Second)
	} */

	// for com range. Primeiro paramentro é o indice e o outro o item
	nomes := [3]string{"Emerson", "Marta", "Stella"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	// Omitindo o índice
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	// Interando sobre palavras. Usa-se a função string() para converter o código ASC para letra
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	// Interando sobre map
	usuario := map[string]string{
		"nome":      "Emerson",
		"sobrenome": "Leite",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// Loop infinito
	for {
		fmt.Println("Execuntando infinitamente")
	}
}
