package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Emerson Batista")
	escrever("Leite")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
