package main

import (
	"fmt"
	"time"
)

func main() {

	// Definindo os canais
	canal1, canal2 := make(chan string), make(chan string)

	// Mandando do canal 1
	go func() {
		for /* x := 0; x < 20; x++ */ {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
		/* close(canal1) */
	}()

	// Mandando do canal 2
	go func() {
		for /* x := 0; x < 5; x++ */ {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
		/* close(canal2) */
	}()

	// Select evita perda de tempo entre sincronizações entre loops,
	// uma vez que o tempo de envio entre os canais é diferente
	for {
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}
	}
}
