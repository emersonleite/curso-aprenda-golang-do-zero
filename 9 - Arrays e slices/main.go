package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Array com quantidade de elementos e mesmo tipo de dado, nesse caso int. Sem inicializar valores. Todos são valores ''
	var array1 [5]string
	array1[0] = "Emerson"
	fmt.Println(array1)

	// Outra forma de criar array inicializando valores
	array3 := [5]string{"Emerson", "Marta", "Stella", "Mike", "Mariana"}
	fmt.Println(array3)

	// Erro
	// array3[5] = "Ola"

	// Array com a quantidade de elementos de acordo com a inicialização. Não seria permitido adicionar além de 5 elementos
	array2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	// Slices não tem tamanho fixo. Slice não é um array
	slice := []int{10, 11, 12, 13}
	fmt.Println(slice)

	// Mostrando os tipos
	fmt.Println(reflect.TypeOf(slice))  // []int
	fmt.Println(reflect.TypeOf(array3)) // [5]int

	// adicionando item a um slice
	slice = append(slice, 10)
	fmt.Println(slice)

	// Criando um slice a partir de valores de um array, sendo que inclui indice 1 e exclui indice 3 (até 2). A referência do slice é por ponteiro. Se mudar valor em array2, muda no slice2
	slice2 := array2[1:3]
	fmt.Println(slice2)

	// ARRAYS INTERNOS
	// make(tipo, tamanho, capacidade máxima)
	fmt.Println("-------------")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	fmt.Println(len(slice3)) // Mostrando o tamanho do slice
	fmt.Println(cap(slice3)) // Mostrando a capacidade do slice

	slice3 = append(slice3, 15)
	slice3 = append(slice3, 12)

	fmt.Println(len(slice3))
	fmt.Println(cap(slice3)) // Dobra a capacidade quando é excedida

	// Quando não se passa a capacidade ele fica sendo igual ao tamanho, ou seja, 5 conforme abiaxo
	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4)) // 6
	fmt.Println(cap(slice4)) // 12

	fmt.Println(slice4)

}
