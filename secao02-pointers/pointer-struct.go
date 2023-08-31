package main

import "fmt"

type Empregado struct {
	Name string
	id   int
}

func main() {
	emp := Empregado{
		Name: "Felippe",
		id:   123,
	}

	pts := &emp

	fmt.Println(pts)
	pts.Name = "Mozer"
	fmt.Println("Valor ponteiro:", pts)
}
