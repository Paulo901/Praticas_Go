package exercicios

import "fmt"

type veiculos struct {
	cor    string
	portas int
}

type caminhonete struct {
	categoria   veiculos
	quatrorodas bool
}
type sedan struct {
	categoria  veiculos
	modeloluxo bool
}

func Exer3() {

	objeto := caminhonete{categoria: veiculos{"preto", 4}, quatrorodas: true}
	objeto2 := sedan{categoria: veiculos{"verde", 4}, modeloluxo: false}

	fmt.Println(objeto, "\n \n", objeto2)
	fmt.Println(objeto.categoria.cor, "\n \n", objeto2.categoria.portas)

}
