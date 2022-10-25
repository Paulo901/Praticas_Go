package exercicios

import "fmt"

func Chamado() {
	fmt.Println("mensagem")
}

func Exer9(f func()) {
	fmt.Println("para não se perder: ")
	f()
}
