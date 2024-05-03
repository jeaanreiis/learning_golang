package main

import "fmt"

func main() {

	// ordem de precedencia de operações

	// '()' -> qualquer conta assim como na matemática é realizada primeino no parenteses
	var x int = 10 * 10 // multiplicação
	var y int = 10 / 10 // divisão
	var z int = 35 % 3  // resto da divisão

	var a int = 5 + 2   // adição
	var b int = 10 - 20 // subtração

	fmt.Println(x, y, z, a, b)
}
