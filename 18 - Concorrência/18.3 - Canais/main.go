package main

import (
	"fmt"
	"time"
)

func main() {

	// Definindo um canal do tipo string
	canal := make(chan string)

	go escrever("Olá mundo", canal)

	// Canal esperando receber o valor.  linha abaixo só é executada quando mensagem recebe o valor do canal, Há também a verificação se o canal recebe algum valor. Se não receber, acontece o break e o laço for é abandonado.

	/* for {
		mensagem, aberto := <-canal
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	} */

	// Outra opção abaixo, sem uso de 'aberto' e 'if':
	for mensagem := range canal {
		fmt.Println(mensagem)
	}
	fmt.Println("Fim do programa")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		// O canal recebe a variável do tipo string
		canal <- texto
		time.Sleep(time.Second)
	}
	// A função close fecha o canal
	close(canal)
}
