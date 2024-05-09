package main

import (
	"fmt"
)

func main() {
	x := [5]float64{98, 93, 77, 82, 83} // para criar um array (array tem tamanho fixo) é nescessario especificar o tipo, mesmo na declaracao mais simples

	var total float64 = 0

	for _, value := range x { // foi usado "_" no lugar do "i" porque o compilador de go nao aceita variaveis nao utilizadas
		total += value // 'value' usado no lugar do 'x[i]'
	}

	fmt.Println(total / float64(len(x)))

}
