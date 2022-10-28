package exercicios

import "fmt"

type Pessoa struct {
	fala string
}

func (p *Pessoa) Falar() {
	p.fala = "esta é minha fala"
}

type Humano interface {
	Falar()
}

func DizendoAlgo(h Humano) {
	h.Falar()
}

func Exer2() {

	p1 := Pessoa{}
	DizendoAlgo(&p1)
	//DizendoAlgo(p1) <- dá erro
	fmt.Println(p1)
}
