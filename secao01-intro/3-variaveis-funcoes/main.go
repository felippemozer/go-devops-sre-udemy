package main

import "fmt"

func main() {
	// Declaração padrão de variável sem atribuição
	// var name string

	// Declaração de variável com atribuição
	name := getName()
	idade := 26
	fmt.Println("Hello", name, idade)
}

func getName() string {
	return "Felippe"
}
