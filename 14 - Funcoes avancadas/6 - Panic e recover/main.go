package main

import "fmt"

func recuperarExecucao() {
	// fmt.Println("Tentativa de recuperação")
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso")
	}
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	}
	if media < 6 {
		return false
	}

	// Em caso de média ser exatamente 6. Mata a execução do programa. Chama antes as funções com 'defer'
	panic("A média é exatamente 6")
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))

	// Com panic (sem recover) essa linha não é executada
	fmt.Println("Pós execução")
}
