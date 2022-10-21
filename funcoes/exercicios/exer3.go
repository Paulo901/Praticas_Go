package exercicios

import "fmt"

func Defer() {

	fmt.Println("Bom dia")
	defer fmt.Println("Boa noite")
	defer fmt.Println("Boa tarde")
}
