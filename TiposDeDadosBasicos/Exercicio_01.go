package main

import (
	"fmt"
)

// 1) Descubra por que não compila
// 2) O que o erro ./prog.go:9:14: constant 150 overflows int8 nos indica?
// 3) Conserte o erro de compilação e faça a quantidade de quilômetros ser imprimida na tela

func main() {
	var quilometros int16
	quilometros = 150

	fmt.Println(quilometros)
}
