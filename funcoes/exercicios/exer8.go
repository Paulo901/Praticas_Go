package exercicios

import "fmt"

func Exer8() func() {

	return func() {
		fmt.Println("teste")
	}
}
