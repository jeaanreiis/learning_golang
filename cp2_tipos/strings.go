package main

import "fmt"

func main() {

	var texto string = "Hello, world"

	fmt.Println(texto)

	fmt.Println("tamanho do texto:", len(texto)) // imprime o tamanho dos caracters

	fmt.Println(texto[1]) // imprime valor 101 que corresponde a letra 'e' seguindo tabela ASCII

	fmt.Println("My Text: " + texto) // o sinal de + também concatena as strings

}
