package main

import (
	"fmt"
	"time"
)

func main() {

	// Início da execução
	start := time.Now()

	fmt.Println("Começou")

	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}

	fmt.Println("acabou")

	// Final da execução
	elapsed := time.Since(start)
	fmt.Printf("Durou %s", elapsed)

}

// '<-chan' canal que só envia dados, 'chan<-' canal que só recebe dados
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
