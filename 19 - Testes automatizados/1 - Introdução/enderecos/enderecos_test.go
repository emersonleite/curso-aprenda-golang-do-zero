// teste de unidade
// Podem estar em outro pacote (outro nome), nesse caso com o nome enderecos_test
// O nome do arquivo deve terminar em '_test'
// O nome da função de teste não precisa ter algum nome em formato especial, mas por boa prática usa-se o nome da função a testar com o pefixo 'Test', como, por exemplo : 'TestObterTipoDeEndereco'
// Para rodar os testes em todas as pastas 'go test ./...'
// Para detalhes 'go test -v' - modo verboso
// Para verificar cobertura 'go test -cover'
// Gerar relatório para verificar cobertura 'go test --coverprofile resultado.txt'
// Verificar com precisão o arquivo de resultado 'go tool cover --func=resultado.txt'
// Verificar em ht,l com precisão o arquivo de resultado 'go tool cover --html=resultado.txt'
// Para rodar testes em paralelo 't.Parallel()'
// Na importação do pacote pode se colocar '.' para um aliás para a função

package enderecos_test

import (
	. "testes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoIserido string
	retornoEsperado string
}

/* func TestObterTipoDeEndereco(t *testing.T) {
	enderecoParaTeste := "Avenida dos Cornos"
	tipoDeEnderecoEsperado := "Rua"
	tipoDeEnderecoRecebido := ObterTipoDeEndereco(enderecoParaTeste)

	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		t.Error("O tipo recebido é diferente do esperado")
		// Para formatar a mensagem
		t.Errorf("Esperava %s e recebeu %s", tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
	}
} */

func TestObterTipoDeEnderecoRefatorada(t *testing.T) {
	// Para rodar testes em paralelo
	t.Parallel()

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		// {"Rodovia Ruela", "Rodovia"},
		{"Estrada da alegria", "Estrada"},
		// {"Prala das raposas", "Tipo inválido"},
		{"RUA DOS BOBOS", "Rua"},
		// {"", "Tipo inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := ObterTipoDeEndereco(cenario.enderecoIserido)

		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s", tipoDeEnderecoRecebido, cenario.retornoEsperado)
		}
	}
}

func TestAnyAction(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Error("Teste quebrou")
	}
}
