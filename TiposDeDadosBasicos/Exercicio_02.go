package main

import (
	"fmt"
	"time"
)

// 1) Dado o ano em que a pessoa nasceu, calcule quantos anos ela tem no ano atual (para fins de arredondamento, considere que ela já fez aniversário no ano atual).
// 2) Como podemos pegar a informação do ano diretamente do console.

func main() {
	fmt.Println("Digite o ano que vc nasceu:")

	var birthYear int
	fmt.Scan(&birthYear)

	todayYear := time.Now().Year()

	age := todayYear - birthYear

	fmt.Println(age)

}
