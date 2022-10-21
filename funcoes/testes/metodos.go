package testes

type Valor struct {
	Num int
}

func (n *Valor) Soma(valor int) {

	n.Num += valor

}
