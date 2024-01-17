package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// Mais de um retorno para uma função
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	// Aplicando a função soma
	soma := somar(1, 3)
	fmt.Println(soma)

	// Igualando função à uma variável 'fu'
	var fu = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	// Aplicando função fu
	fu("ola")

	// Igualando a função fu a resultado
	resultado := fu("ola 2")
	fmt.Println(resultado)

	// Aplicando a função calculosMatematicos
	resultadoSoma, resultadoSubtracao := calculosMatematicos(4, 2)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	// Aplicando calculosMatematicos com omissão de um dos retornos
	_, resultadoSubtracao2 := calculosMatematicos(4, 2)
	fmt.Println(resultadoSubtracao2)

}
