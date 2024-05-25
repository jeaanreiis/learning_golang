package main

import (
	"fmt"
)

func main() {
	// forma mais concisa de declarar mapas
	mapa := map[string]int{
		"key": 10,
	}

	fmt.Println(mapa["key"])
	// declaracao 2
	elements := make(map[string]string)
	elements["H"] = "Helio"
	elements["B"] = "Boron"

	name, ok := elements["H"]
	fmt.Println(name, ok) // o 'ok' guarda valor booleano para idenficar se existe o valor no dicionario

	m := map[string]string{
		"H":  "Hidrogênio",
		"He": "Hélio",
		"Li": "Lítio",
		"Be": "Berílio",
		"C":  "Carbono",
		"B":  "Boro",
		"N":  "Nitrogênio",
		"O":  "Oxigênio",
		"F":  "Flúor",
		"Ne": "Neônio",
		"Na": "Sódio",
		"Mg": "Magnésio",
		"Al": "Alumínio",
		"Si": "Silício",
		"P":  "Fósforo",
		"S":  "Enxofre",
		"Cl": "Cloro",
		"Ar": "Argônio",
	}

	fmt.Println(m["Mg"])

	if elemento, ok := m["Rn"]; ok {
		fmt.Println("Elementro encontrado", elemento, ok)
	} else {
		fmt.Println("Elementro NÃO encontrado", elemento, ok)
	}

	maps := map[string]map[string]string{
		"H": map[string]string{
			"nome":   "Hidrogênio",
			"estado": "Gás",
		},
	}

	if el, ok := maps["H"]; ok { // o ok nao precisa ser printado, pode se usado para estrutura de controle
		fmt.Println(el["nome"], el["estado"])
	}

}
