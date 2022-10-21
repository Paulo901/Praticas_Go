package exercicios

import "fmt"

type Pessoa struct {
	Nome, Sobrenome string
	Idade           int
}

func (p *Pessoa) Ola() {
	fmt.Println(p.Nome, p.Sobrenome, ", você tem", p.Idade, "anos")
}
