package main

import (
	"fmt"
)

// Faça um programa que solicite à usuária que informe um número até que o número informado seja par.

func main() {

	fmt.Println("Digite um número:")

	var numero int
	fmt.Scan(&numero)

	for (numero % 2) != 0 {
		fmt.Println("Digite um número:")
		fmt.Scan(&numero)
	}
	fmt.Println("Finalmente um numero par")
}
