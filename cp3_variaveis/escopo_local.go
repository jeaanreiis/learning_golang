package main

import "fmt"

func main() {
	var (
		a = 5
		b = 10
		c = 15
	)

	fmt.Println(a, b, c)

	f() // variavel declarada no escopo local
}

func f() {
	a := 100

	fmt.Println("Minha nova func:", a)
}
