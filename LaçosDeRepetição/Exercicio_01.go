package main

import "fmt"

// No Exercício #06 da seção "Exercícios", usamos for range para percorrer uma slice de string que representava uma lista de itens a comprar no mercado.
// Agora, resolva o mesmo exercício usando a sintaxe básica da instrução for (sintaxe apresentada aqui).

func main() {
	listaDeMercado := []string{"abacate", "sabonete", "azeite de oliva", "tomate", "banana", "macarrão", "cebola"}

	for i := 0; i < len(listaDeMercado); i++ {
		fmt.Printf("%d) %s\n", i+1, listaDeMercado[i])

	}
}
