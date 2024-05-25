// projetos simples para treinar golang //

// criar um programa para calcular notas de alunos (conceito estrutura de controle e decisao)

// criar programa "breaking bad" que devolve informacoes tabela periodica de acordo com letra ou nome do elemento inseriodo no prompt (conceitos de mapas)

// criar calculadora (conceito de funcoes)

package main

import "fmt"

func main() {
	fmt.Printf("Olá, Mundo!")

	m := make(map[string]string)

	m["info"] = "Meu texto"

	fmt.Println(m["info"])

}
