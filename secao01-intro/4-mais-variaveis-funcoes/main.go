package main

import "fmt"

func main() {
	name, salary := "Felippe", 2000
	setName(name)
	newSalary, bonus := addSalaryV2(salary, 100)
	fmt.Println("Novo salario:", newSalary)
	fmt.Println("Bonus:", bonus)
}

// função void com parâmetro
func setName(name string) {
	fmt.Println("Hello", name)
}

// função com parâmetros e retorno simples
func addSalary(valorAtual int, bonus int) int {
	return valorAtual + bonus
}

// função com retorno duplo
func addSalaryV2(valorAtual int, bonus int) (int, int) {
	return valorAtual + bonus, bonus
}
