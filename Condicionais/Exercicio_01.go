package main

import "fmt"

// Faça um programa em que 3 variáveis recebem valores diferentes e informa qual a variável com maior valor.

func main() {

	numero1 := 41
	numero2 := 6
	numero3 := 28

	maiorValor := 0
	nomeVariavel := ""

	if numero1 > maiorValor {
		maiorValor = numero1
		nomeVariavel = "numero1"
	}
	if numero2 > maiorValor {
		maiorValor = numero2
		nomeVariavel = "numero2"
	}
	if numero3 > maiorValor {
		maiorValor = numero3
		nomeVariavel = "numero3"
	}
	fmt.Println(nomeVariavel)

}
