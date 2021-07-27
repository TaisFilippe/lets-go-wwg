package main

import "fmt"

// Faça um programa que testa se um número passado em uma variável é 0, múltiplo de 2, múltiplo de 3 ou nenhuma das opções.

func main() {

	fmt.Println("Digite um número:")

	var numero int
	fmt.Scan(&numero)

	switch {
	case numero == 0:
		fmt.Println("O número é igual a zero")
	case (numero%2) == 0 && (numero%3) == 0:
		fmt.Println("O número é multiplo de 2 e de 3")
	case (numero % 2) == 0:
		fmt.Println("O número é multiplo de 2")
	case (numero % 3) == 0:
		fmt.Println("O número é multiplo de 3")
	default:
		fmt.Println("Nenhuma das opções")
	}
}
