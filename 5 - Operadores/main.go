package main

import "fmt"

func main() {

	// OPERADORES ARITMÉTICOS
	soma := 1 + 2
	subtracao := 2 - 1
	divisao := 1 / 2
	multiplicacao := 2 * 1
	restoDaDivisao := 4 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	// Não é possível somar var com tipo diferentes
	var numero2 int16 = 10
	var numero3 int32 = 20
	// soma1 := numero2 + numero3 ---> erro
	fmt.Println(numero2, numero3)

	// OPERADORES ATRIBUICAO
	var nome string = "Emerson Leite"
	titulo := "dev"
	fmt.Println(nome, titulo)

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// OPERADORES LÓGICOS
	fmt.Println(false && true)
	fmt.Println(false || true)
	fmt.Println(!false)

	// OPERADORES UNÁRIOS
	numero := 10
	numero++
	fmt.Println(numero)

	numero--
	numero--
	fmt.Println(numero)

	numero += 10
	fmt.Println(numero)

	numero -= 8
	fmt.Println(numero)

	numero *= 3
	fmt.Println(numero)

	numero /= 3
	fmt.Println(numero)

	numero %= 3
	fmt.Println(numero)

}
