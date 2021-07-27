package main

import "fmt"

// Escreva um programa que lista o nome dos países e quantas vezes eles aparecem no nosso map.

// Passos:
// 1)Criar um mapa com chave do tipo string e valor do tipo string (map[string]string) onde a chave seja o nome da cidade e o valor o nome do país.
// 2)Percorrer o mapa do item 1 acumulando em outro mapa a frequência de aparições do país. Esse segundo mapa terá tipo map[string]int, sendo a chave o nome do país e o valor a quantidade de menções.
// 3)Printe na tela os valores do último mapa.

func main() {

	meuMapa := make(map[string]string)

	meuMapa["Paranagua"] = "Brasil"
	meuMapa["Lyon"] = "França"
	meuMapa["Curitiba"] = "Brasil"
	meuMapa["Los Angeles"] = "USA"
	meuMapa["Florenca"] = "Italia"
	meuMapa["São Paulo"] = "Brasil"
	meuMapa["Veneza"] = "Italia"

	outroMapa := make(map[string]int)

	for _, pais := range meuMapa {
		quantidade, existe := outroMapa[pais]
		if existe {
			outroMapa[pais] = quantidade + 1
		} else {
			outroMapa[pais] = 1
		}
	}

	fmt.Printf("%#v\n", outroMapa)

}
