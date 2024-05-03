package main

import "fmt"

func main() {

	fmt.Println("\nconjunção também conhecido como 'E' lógico:")

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && true)
	fmt.Println(false && false)
	// exemplo:
	var x bool = true
	var y bool = false
	var z bool = x && y // z será 'false'
	fmt.Println("Resultado com operador 'E': ", z)

	fmt.Println("\ndisjunção também conhecido como 'OU' lógico:")
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(false || false)
	// exemplo:  !!! variaveis declaradas anteriormente !!!
	x = true
	y = false
	z = x || y // z será 'true'
	fmt.Println("Resultado com operador 'OU': ", z)

	fmt.Println("\nnegação também conhecido como 'não igual a' lógico:")
	fmt.Println(!true)
	fmt.Println(!false)
	// exemplo:
	var a bool = true
	var b bool = !a // b será false
	fmt.Println("Resultado com operador 'Diferente de': ", b)

}
