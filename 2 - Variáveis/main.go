package main

import "fmt"

func main() {

	// Com definição do tipo
	var message string = "Olá torto"
	fmt.Println(message)

	// Com inferência de tipo
	outraMensagem := "Olá lindão"
	fmt.Println(outraMensagem)

	// Outra forma
	var (
		message3 string = "Ola"
		message4 string = "mundo"
	)
	fmt.Println(message3, message4)

	// Outra forma, ufa...
	message5, message6 := "Golang", "é legal"
	fmt.Println(message5, message6)

	// Trocando valores de variáveis
	message5, message6 = message6, message5
	fmt.Println(message5, message6)

}
