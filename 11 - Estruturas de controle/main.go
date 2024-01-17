package main

import "fmt"

func main() {

	numero := -11

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	// Inicializando a variável de controle a partir de outra
	// Variável faz parte do escopo do if somente
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que 0")
	} else if outroNumero < -10 {
		fmt.Println("Número é menor que -10")
	} else {
		fmt.Println("Entre 0 e -10")
	}

	// Erro
	// fmt.Println(outroNumero)

}
