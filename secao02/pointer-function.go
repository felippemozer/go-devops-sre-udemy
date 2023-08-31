package main

import "fmt"

func pointerFunc(a *int) {
	*a = 400
}

func main() {
	// Parte 2
	var x = 100
	fmt.Println("O valor da variável X antes da função é:", x)
	var pa *int = &x

	pointerFunc(pa)

	fmt.Println("O valor da variável X depois da função é:", x)
}
