package exercicios

import "fmt"

type pessoa struct {
	nome string
	tipo string
}

func Exer2() {
	alguem := pessoa{nome: "Fulano", tipo: "não estuda"}
	fmt.Println(alguem)
	MudeMe(&alguem)
	fmt.Println(alguem)
}

func MudeMe(p *pessoa) {
	p.tipo = "super estudioso"
}
