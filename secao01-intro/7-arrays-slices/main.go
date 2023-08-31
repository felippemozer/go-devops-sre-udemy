package main

import "fmt"

func main() {
	// Arrays em GO são estáticos
	var salarios [10]int
	salarios[0] = 1000
	salarios[1] = 1500

	// Slices (arrays dinâmicos)
	salarios2 := []int{1000, 1500, 2000, 3000}
	// salarios2 := make([]int, 5)

	for i := 0; i < len(salarios2); i++ {
		salarios2[i] += 100 + i
	}

	for _, salario := range salarios2 {
		fmt.Println(salario)
	}

	salarios3 := []int{}

	for i := 0; i < 10; i++ {
		salarios3 = append(salarios3, 100+i)
	}

	fmt.Printf("salarios3: %v\n", salarios3)
}
