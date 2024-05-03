package main

import "fmt"

var (
	a = "Red"
	b = "Green"
	c = "Blue"
)

func main() {
	fmt.Println(a, b, c)

	f() // minha funcao para escopo global
}

func f() {
	fmt.Println("Minha nova func: "+a, b, c)
}
