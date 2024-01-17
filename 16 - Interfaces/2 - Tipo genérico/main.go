package main

import "fmt"

// Parâmetro do tipo interface genérica
func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica(1)
	generica("Ola")
	generica(true)

}
