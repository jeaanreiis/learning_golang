package main

import "fmt"

func main() {

	var a int
	var b, c string // posso declarar outras variaveis do mesmo tipo
	var x bool
	var y float64

	a = 10
	b = "hello"
	c = "Declaração inicializada separadamente: "
	x = !true
	y = 3.14

	fmt.Println(c, a, b, x, y)

}
