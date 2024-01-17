package main

import (
	"fmt"
)

func main() {

	// Passagem de valor por cópia, ou seja, se muda o valor de v1 não muda o valor de v2
	var variavel1 int = 10
	var variavel2 int = variavel1
	fmt.Println(variavel1, variavel2)

	// Só se muda o valor de v1, v2 permace inalterado
	variavel1++
	fmt.Println(variavel1, variavel2)

	// Ponteiro é uma referência de memória

	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	// Associa ao ponteiro o endereço de memória onde o valor 100 está salvo
	ponteiro = &variavel3

	fmt.Println(variavel3)

	// Aqui se mostra o valor do endereço de memória de ponteiro
	fmt.Println(ponteiro)

	// Aqui se faz a desreferenciação para mostrar o valor númerico do ponteiro
	fmt.Println(*ponteiro)

	variavel3 = 150

	// Mostrando o valor atualizado do ponteiro
	fmt.Println(*ponteiro)

}
