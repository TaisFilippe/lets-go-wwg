package main

// Construa uma função que receba uma palavra ou frase e uma letra, e retorne o número de ocorrências da letra informada.

import (
	"fmt"
)

func main() {
	frase := leiaLinha()
	ocorrencias := make(map[byte]int)

	for i := 0; i < len(frase); i++ {

		quantidade, existe := ocorrencias[frase[i]]
		if existe {
			ocorrencias[frase[i]] = quantidade + 1
		} else {
			ocorrencias[frase[i]] = 1
		}
	}

	for letra, valor := range ocorrencias {
		fmt.Printf("%c: %v \n", letra, valor)
	}
}

func leiaLinha() string {
	return "Women Who Go Curitiba"

}
