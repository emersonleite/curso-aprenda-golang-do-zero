package main

import "fmt"

func main() {
	fmt.Println("Função main sendo executada")
}

// Função init é executada antes da função main.
// Pode ter uma função por arquivo, ao contrário da main que é única por pacote.
func init() {
	fmt.Println("Executando a função init")
}
