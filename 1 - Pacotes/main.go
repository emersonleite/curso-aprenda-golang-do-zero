package main

import (
	"fmt"
	"modulo_teste/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")

	auxiliar.Mostrar()

	erro := checkmail.ValidateFormat("emerson.bl@gmail.com")

	erro2 := checkmail.ValidateFormat("emerson.blgmail.com")

	fmt.Println(erro) // saida -> '<nil>'

	fmt.Println(erro2) // saida -> 'invalid format'

}
