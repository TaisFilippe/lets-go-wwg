package main

import "fmt"

//Considerando os tópicos que já aprendemos até agora: slices, structs ,condicionais e laços de repetição, crie um programa que traga as informações sobre apartamentos de um prédio.

// Passos:
// Crie uma estrutura que representa um apartamento, com campos para representar seu número, o nome da sua proprietária e se tem vaga de garagem
// Reúna as estruturas em uma slice que representa um conjunto de apartamentos
// Printe as informações de cada unidade, separando por linha, usando for range

type Apartamento struct {
	numero              int
	nomeDaProprietaria  string
	possuiVagaDeGaragem bool
}

func main() {
	ap1 := Apartamento{
		numero:              1102,
		nomeDaProprietaria:  "Tais",
		possuiVagaDeGaragem: false,
	}
	ap2 := Apartamento{
		numero:              401,
		nomeDaProprietaria:  "Ana",
		possuiVagaDeGaragem: true,
	}
	ap3 := Apartamento{
		numero:              1003,
		nomeDaProprietaria:  "Eliza",
		possuiVagaDeGaragem: true,
	}
	apartamentos := []Apartamento{ap1, ap2, ap3}

	for _, apartamento := range apartamentos {
		possuiVaga := "Não"
		if apartamento.possuiVagaDeGaragem {
			possuiVaga = "Sim"

		}

		fmt.Printf("O apartamento número %d tem como proprietária %s. Ele tem vaga de garagem? %s\n", apartamento.numero, apartamento.nomeDaProprietaria, possuiVaga)
	}
}
