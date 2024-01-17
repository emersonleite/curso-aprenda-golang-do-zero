package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Println("Dentro do método salvar")
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

// Método com passagem de parâmetro por referência com o '*' antes do tipo
func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {

	const PI = 3.14
	fmt.Println(PI)

	usuario1 := usuario{
		nome:  "Emerson",
		idade: 41,
	}

	fmt.Println(usuario1)

	usuario1.salvar()
	maiorDeIdade := usuario1.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	// Altera de fato a idade do usuário, por a passagem de parâmetro é por referência
	usuario1.fazerAniversario()
	fmt.Println(usuario1.idade)

}
