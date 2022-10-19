package exercicios

import "fmt"

var a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

func Exer2() {

	fmt.Printf("%T \n", a)

	for _, value := range a {

		fmt.Println(value)
	}

}
