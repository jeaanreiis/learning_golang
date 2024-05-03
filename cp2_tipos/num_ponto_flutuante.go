package main

import "fmt"

func main() {

	//em go existe tipo float32 para precisao simples e float64 para precisao dupla

	const pi float64 = 3.14159265358979323846264338327950288419719399375105820974
	const e float64 = 2.71828

	fmt.Printf("%.2f\n", pi) // formatacao de casas decimais com quebra de linha

	e_fmt := fmt.Sprintf("%.2f", e) // formatacao usando sprintf para tipo de dado string

	fmt.Println(e_fmt) // saída
}
