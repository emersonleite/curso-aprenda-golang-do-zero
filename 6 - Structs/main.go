package main

import "fmt"

type usuario struct {
	nome     string
	Idade    uint8
	endereco endereco
}

type endereco struct {
	rua    string
	numero uint16
}

func main() {

	enderecoExemplo := endereco{"Fernandes de Barros", 1732}

	// Definindo na inicialização
	// emerson := usuario{nome: "Emerson", Idade: 41} // ou
	emerson := usuario{"Emerson", 41, enderecoExemplo}
	fmt.Println(emerson)
	fmt.Println(emerson.endereco)

	// Definindo depois de inicializar
	var mariana usuario
	mariana.nome = "Mariana"
	mariana.Idade = 19
	fmt.Println(mariana)

	// Sem definir todos os dados de uma vez
	stella := usuario{nome: "Stella"}
	fmt.Println(stella)

}
