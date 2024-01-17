package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) string {
	// 'defer' adia a execução até o último momento possível. No caso de funções com retorno, até o retorno.
	defer fmt.Println("Média calculada. Resultado será retornado. Essa linha utiliza defer e será mostrada antes do retorno somente")

	fmt.Println("Entrando na função para verificar se o aluno está aprovado")

	media := (n1 + n2) / 2

	var text string = "REPROVADO"

	if media >= 7.0 {
		text = "APROVADO"
	}
	return text
}

func main() {

	// Adiando a execução da função 1

	/* defer funcao1()
	funcao2() */

	fmt.Println(alunoEstaAprovado(6, 5))

}
