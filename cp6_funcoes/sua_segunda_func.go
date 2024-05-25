package main

import "fmt"

func calculo_media(notas []float64) float64 { // float64 é referente ao retorno da funcao
	// uma função também é conhecida como procedimento ou sub-rotina

	total := 0.0

	for _, nota := range notas {

		total += nota
	}

	return total / float64(len(notas))
}

func f2() (x int) { // outra maneira de retornar a funcao
	x = 10

	return
}

// retorno de multiplos valores
func valores() (int, int) {
	return 1, 2
}

// funcao principal
func main() {

	notas := []float64{10, 7.7, 8.7, 8.3, 9.0}

	fmt.Println(calculo_media(notas))

	fmt.Println(f2())

	x, y := valores()

	fmt.Print(x, y)
}
