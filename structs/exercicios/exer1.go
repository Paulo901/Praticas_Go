package exercicios

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	hobby     []string
}

func Exer1() {

	p1 := []pessoa{
		{nome: "Paulo", sobrenome: "Oliveira", hobby: []string{"Xadrez"}},
		{nome: "Outro", sobrenome: "Fulano", hobby: []string{"natação"}},
	}

	fmt.Println(p1)
	for _, value := range p1 {
		fmt.Println(value.hobby)
	}
}
