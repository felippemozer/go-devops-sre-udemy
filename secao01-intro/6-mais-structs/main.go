package main

import "fmt"

type Pessoa struct {
	name   string
	age    int
	salary int
}

func main() {
	// Declaração inicial de struct com atribuição inicial
	p := &Pessoa{
		name:   "Felippe",
		age:    26,
		salary: 2000,
	}
	p.addSalary(100)
	fmt.Println(p.salary)

	// Declaração inicial de struct sem atribuição inicial
	p2 := new(Pessoa)
	p2.name = "Pessoa2"
	p2.age = 30
	p2.salary = 1000
}

// Referência de função a um struct
func (p *Pessoa) addSalary(bonus int) {
	p.salary += bonus
}
