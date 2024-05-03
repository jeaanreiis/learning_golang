package main

import (
	"fmt"
)

func main() {

	// EXEMPLO DE LACO 'FOR'
	for i := 1; i <= 10; i++ { // inicializacao ( i:= N); condicao (i <= N); incremento (i++)

		fmt.Println(i, "primeira contagem do loop...")

	}

	fmt.Println("Fim do primeiro loop!")

	fmt.Println("\n") // quebrar linha com espaco maior

	// OUTRO MODELO DE DECLARACAO DE CONDICIONAL 'FOR'
	i := 10

	for i >= 1 {
		fmt.Println(i, "segunda contagem do loop...")

		i-- // inverso do incremento (i++)
	}

	fmt.Println("Fim do segundo loop!")

	fmt.Println("\n") // quebrar linha com espaco maior

}
