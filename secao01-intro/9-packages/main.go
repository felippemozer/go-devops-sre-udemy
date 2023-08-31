package main

import (
	"fmt"
	"packages/funcionarios"
)

func main() {
	fmt.Println("Bem vindo, Felippe")
	p := funcionarios.Pessoa{
		name: "Felippe",
	}

	fmt.Printf("p: %v\n", p)
}
