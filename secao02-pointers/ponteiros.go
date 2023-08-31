package main

import "fmt"

func main() {
	// Parte 1
	var inteiro = 45
	var ponteiro *int = &inteiro

	fmt.Println("Valor da variável inteira:", inteiro)
	fmt.Println("Endereço da variávei inteiro:", &inteiro)
	fmt.Println("Valor da variável ponteiro:", ponteiro)
	fmt.Println("Endereço da variável ponteiro:", &ponteiro)
	fmt.Println("Valor da variável inteiro pelo ponteiro:", *ponteiro)

}
