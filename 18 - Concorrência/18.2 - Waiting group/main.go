package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup

	// Adicionando 2 goroutines que rodarão ao mesmo tempo
	waitGroup.Add(2)

	go func() {
		escrever("Emerson Batista")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Leite")
		waitGroup.Done() // -1
	}()

	// Evitar que o programa termine antes da hora
	waitGroup.Wait()

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
