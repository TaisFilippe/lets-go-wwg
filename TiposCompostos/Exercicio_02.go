package main

import "fmt"

// Considerando os times do item anterior, crie uma slice para representar cada um.
// Adicione o jogador Luis Inácio no time vermelho usando a função append().
// Printe na tela os nomes dos jogadores do time vermelho .

func main() {

	//TimeAmarelo := []string{"Fernando", "João", "Lúcia", "Mariana", "Ana"}

	TimeVermelho := []string{"Helena", "Jonas", "José", "Juliana"}
	TimeVermelho = append(TimeVermelho, "Luis Inácio")

	//fmt.Printf("Jogadores do Time Amarelo: %s\n", TimeAmarelo)

	fmt.Printf("Jogadores do Time Vermelho: %s\n", TimeVermelho)

}
