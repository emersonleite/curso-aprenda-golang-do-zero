# Pacotes

- Criação de módulos, que são conjunto de pacotes

```shell
go mod init nome_do_modulo

```

- Para dar build no projeto

```shell
go build

```

- Para se criar um pacote, se cria uma pasta. O nome do arquivo do pacote pode ser qualquer.
  O que o diferencia é o header do arquivo

- Nome do pacote no header com arquivo:

```go
package auxiliar
```

- Para que elementos fiquem disponível do módulo (sendo públicas) é necessário que iniciem com letra maiúscula. Necessário comentar funções exportadas;

```go
package auxiliar

import "fmt"

var t int

// Função mostrar
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

```

# Pacotes externos

- Obtendo:

```shell
go get github.com/badoux/checkmail

```

- Retirando pacotes não utilizados no projeto:

```shell
go mod tidy
```

- Para adicionar pacote ao workspace:

```shell
go work use .
```
