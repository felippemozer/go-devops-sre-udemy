package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Tratamento de erros
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("erro, nao é um número válido")
		os.Exit(1)
	}
	fmt.Println("Numero convertido:", n)

	// Defer
	file, err := os.Open("file.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
}
