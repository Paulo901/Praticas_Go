package exercicios

import "fmt"

func Exer6() {

	a := make([]string, 1, 27)
	fmt.Println(len(a), cap(a))

	a = []string{"Ceará"}
	fmt.Println(a)

}
