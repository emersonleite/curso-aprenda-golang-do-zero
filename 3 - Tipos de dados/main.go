package main

import (
	"errors"
	"fmt"
)

func main() {

	// TIPO int
	var numeroInt64 int64 = 9223372036854775807
	var numeroInt32 int32 = 2147483647
	var numeroInt16 int16 = 32767
	var numeroInt8 int8 = 127

	// Alias para int32
	var numeroRune rune = 2147483647

	// int utiliza a arquitetura do computador como base
	var numeroInt int = 10
	fmt.Println(numeroInt64, numeroInt32, numeroInt16, numeroInt8, numeroRune, numeroInt)

	// TIPO uint -> sem sinal.
	// Seguem os mesmos tamanho de int
	// uint32 ->  Range: 0 through 4294967295
	var numeroSemSinal32 uint32 = 4294967295
	fmt.Println(numeroSemSinal32)

	// TIPO float
	var numeroFloat32 float32 = 123.45
	fmt.Println(numeroFloat32)

	var numeroFloat64 float64 = 100000000000.333333333
	var numeroFloat64_2 = 1.0e+10
	fmt.Println(numeroFloat64_2)
	fmt.Println(numeroFloat64)

	// Não se pode declara tipo float direto
	// var numeroComErro float = 10.00

	// TIPO string, com aspas duplas
	var nome string = "Emerson Leite"
	fmt.Println(nome)

	// TIPO 'char' -> retorna o número na tabela ASC. Não é um tipo próprio, mas sim um int
	char := 'E'
	fmt.Println(char)

	// TIPO boolean
	var isTrue bool = true
	var isFalse bool = false

	fmt.Println(isFalse, isTrue)

	// TIPO error
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

	// Valores sem inicialização
	var (
		n1 int     // 0
		f1 float32 // 0
		t1 string  // ""
		b1 bool    // false
		e1 error   // <nil>
	)
	fmt.Println(n1, f1, t1, b1, e1)

}
