package main

import (
	"fmt"
	"time"
)

func main() {
	canal := createChannel("Ola mundo")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func createChannel(text string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return canal
}
