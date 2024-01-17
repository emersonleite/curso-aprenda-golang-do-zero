package main

import "fmt"

func main() {

	// Definindo canal convencional
	/* canal := make(chan string) */

	// Vai dar erro de deadlock. Necessário funções separadas
	/* mensagem := <-canal */
	// fmt.Println(mensagem)

	// Definindo canal com 2 buffer
	// Canal com buffer só bloqueia se atingir a capacidade máxima.
	canal := make(chan string, 2)
	canal <- "Olá mundo!\n"
	canal <- "Programando em Go!"
	// canal <- "Terceiro valor em Go!"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem, mensagem2)

}
