package auxiliar

import "fmt"

var t int

// Função mostrar que fica disponível fora pois começa com letra maiúscula
func Mostrar() {
	t = 10
	fmt.Println("Escrevendo do módulo auxiliar")
	fmt.Printf("Ola, o número a ser impresso é o %d \n", t)
	mostrar2()
}

// Função interna. Não disponível externamente
func mostrar2() {
	fmt.Println("Ola de dentro de uma função interna do módulo")
}
