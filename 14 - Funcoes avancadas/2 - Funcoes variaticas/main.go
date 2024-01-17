package main

import "fmt"

func somaVariosNumero(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}
	return total
}

// Não pode ter mais de um parâmetro variático. E ele deve esta na ordem correta
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	soma := somaVariosNumero(1, 3, 3, 47, 5, 6, 9, 75, 258)
	fmt.Println(soma)

	escrever("Ola", 1, 5, 9, 8, 5)
}
