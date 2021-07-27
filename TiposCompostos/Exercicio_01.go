package main

import "fmt"

// Existem dois times de futebol, o time amarelo e o time vermelho.
// O time amarelo tem 5 jogadores (Fernando, João, Lúcia, Mariana e Ana) e o time vermelho tem 4 jogadores (Helena, Jonas, José e Juliana).
// Crie um array de string para cada time e nomeie com o nome do time.
// Printe na tela os nomes dos jogadores de cada time

func main() {

	var TimeAmarelo [5]string
	TimeAmarelo[0] = "Fernando,"
	TimeAmarelo[1] = "João,"
	TimeAmarelo[2] = "Lúcia,"
	TimeAmarelo[3] = "Mariana,"
	TimeAmarelo[4] = "Ana"

	var TimeVermelho [4]string
	TimeVermelho[0] = "Helena,"
	TimeVermelho[1] = "Jonas,"
	TimeVermelho[2] = "José,"
	TimeVermelho[3] = "Juliana"

	fmt.Printf("Jogadores do Time Amarelo: %s\n", TimeAmarelo)

	fmt.Printf("Jogadores do Time Vermelho: %s\n", TimeVermelho)

}
