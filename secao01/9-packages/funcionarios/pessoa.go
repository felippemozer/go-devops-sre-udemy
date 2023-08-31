package funcionarios

type Pessoa struct {
	name   string
	age    int
	salary int
}

func (p *Pessoa) addSalary(bonus int) {
	p.salary += bonus
}
