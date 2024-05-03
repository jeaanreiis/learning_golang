package main

import (
	"fmt"
)

func main() {

	// ESTRUTURA DE CONTROLE FOR COM IF
	for i := 1; i <= 10; i++ {

		if i%2 == 0 {
			fmt.Println(i, "PAR")
		} else {
			fmt.Println(i, "ÍMPAR")
		}

	}

	fmt.Println("\n") // quebrar linha com espaco maior

	// ESTRUTURA DE CONTROLE COM OUTRAS CONDICOES
	for i := 10; i >= 0; i-- {
		if i == 10 {
			fmt.Println(i, "...")

		} else if i == 5 {
			fmt.Println(i, "...")

		} else if i == 1 {
			fmt.Println(i, "...")

		}
	}
}
