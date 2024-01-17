package main

import "fmt"

// Passagem de parâmetro por cópia
func inverterSinal(numero int) int {
	return numero * -1
}

// Passagem de parâmetro por referência
func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {

	numero := 20

	numeroInvertido := inverterSinal(numero)

	fmt.Println(numeroInvertido)

	fmt.Println(numero)

	// Aplicação do conceito de ponteiro
	novoNumero := 40
	fmt.Println(novoNumero)

	// Passagem de valor por referência com '&'
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)

}
