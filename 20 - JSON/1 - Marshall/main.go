package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// Para ocorrer o parse correto do JSON
type Cachorro struct {
	Nome  string `json: "nome"`
	Raca  string `json: "raca"`
	Idade uint   `json: "idade"`
}

func main() {

	// A partir de um struct
	c1 := Cachorro{"Rex", "Vira-latas", 3}
	cachorroEmJSON, erro := json.Marshal(c1)
	if erro != nil {
		log.Fatal(erro)
	}
	convertidoEmJSON := bytes.NewBuffer(cachorroEmJSON)
	fmt.Println(convertidoEmJSON)

	// A partir de um map
	c2 := map[string]string{
		"nome": "Totó",
		"raca": "poodle",
	}
	cachorro2EmJSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}
	convertido2EmJSON := bytes.NewBuffer(cachorro2EmJSON)
	fmt.Println(convertido2EmJSON)

}
