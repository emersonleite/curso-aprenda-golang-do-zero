package main

import (
	"fmt"
	"testes/enderecos"
)

func main() {

	tipoDeEndereco := enderecos.ObterTipoDeEndereco("Rua dos Chuletas")

	fmt.Println(tipoDeEndereco)

}
