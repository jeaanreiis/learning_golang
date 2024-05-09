package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3} // alocacao dinamica sem funcao make
	fmt.Println("slice 1: ", slice1)
	slice2 := make([]int, 2) // utilizando funcao make
	fmt.Println("slice 2: ", slice2)
	slice2[0] = 7
	slice2[1] = 8
	fmt.Println("inserindo dados no slice 2: ", slice2)

	slice1 = append(slice1, 5, 6)
	fmt.Println("adicionando dados no slice 1: ", slice1)

	copy(slice2, slice1) // slice 1 foi copiado para slice 2, porém apenas os dois primeiros numero por falta de limite no slice 2

	fmt.Println("copiando informacoes do slice 1 para slice 2: ", slice1, slice2)
}
