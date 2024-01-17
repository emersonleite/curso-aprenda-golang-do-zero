package main

import "fmt"

type Pessoa struct {
	nome string
}

func main() {

	// Chaves tem que ter o mesmo tipo, assim como os valores
	usuario := map[string]string{
		"nome":      "Emerson",
		"sobrenome": "Leite",
	}
	fmt.Println(usuario)

	//  O Acesso não se dá por ponto, mas por colchetes
	fmt.Println(usuario["nome"])

	// Outro exemplo

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Engenharia",
		},
	}
	fmt.Println(usuario2)

	// Para deletar chaves
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	// Para adicionar
	usuario2["signo"] = map[string]string{"nome": "Gemeos"}
	fmt.Println(usuario2)

	// Map considerando uma struct
	mapDeStruct := map[string]Pessoa{
		"p1": {nome: "Emeron"},
	}
	fmt.Println(mapDeStruct)

}
