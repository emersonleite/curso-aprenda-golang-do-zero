package main

import "fmt"

func closureExample() func() {
	texto := "Dentro da função closureExample"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Dentro da função main"
	fmt.Println(texto)

	funcaoNova := closureExample()
	funcaoNova()

}
