package main

import "fmt"

type Pessoa struct {
	name   string
	age    int
	salary int
}

func main() {
	// Inicialização de struct
	pessoa := &Pessoa{
		name:   "Felippe",
		age:    26,
		salary: 2000,
	}
	addSalary(pessoa, 100)
	fmt.Println(pessoa.salary)
}

// Função com ponteiros
func addSalary(p *Pessoa, bonus int) {
	p.salary += bonus
}
