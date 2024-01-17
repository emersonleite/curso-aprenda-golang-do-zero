package enderecos

import "strings"

// Verifica tipo de endereço válido e o retorna
func ObterTipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoEmMinuscula := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmMinuscula, " ")[0]

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			return strings.Title(primeiraPalavraDoEndereco)
		}
	}
	return "Tipo inválido"
}
