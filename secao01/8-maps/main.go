package main

import "fmt"

func main() {
	salFunc := make(map[string]int)
	salFunc["Felippe"] = 10
	salFunc["Luis"] = 20

	sal, exists := salFunc["Luis"]
	fmt.Println(sal, exists)

	totalSal := len(salFunc)
	fmt.Println(totalSal)
}
