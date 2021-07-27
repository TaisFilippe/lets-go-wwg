package main

// Faça um programa que, dado um texto inserido pelo usuário, itere nesse texto e conte o número de ocorrências de cada letra.
// Em seguida imprima em ordem alfabética a letra e o número de ocorrências dela no texto informado.

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
	return "esse texto que voce le agora e o texto informado pelo usuario"
	// sc := bufio.NewScanner(os.Stdin)
	// sc.Scan()
	// return sc.Text()
}
