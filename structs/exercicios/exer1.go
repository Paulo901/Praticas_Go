package exercicios

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
}

func Exer1() {

	p1 := []pessoa{
		{nome: "Paulo", sobrenome: "Oliveira"},
		{nome: "Outro", sobrenome: "Fulano"},
	}

	fmt.Println(p1)
	for _, value := range p1 {
		fmt.Println(value.sobrenome)
	}
}
