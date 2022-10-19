package exercicios

import "fmt"

type Tipo int

var d Tipo

func Exer4() {

	fmt.Println(d)
	fmt.Printf("%T \n", d)
	d = 42
	fmt.Println(d)

}
