package testes

import "fmt"

func Funcoes() {

	tes := []int{1, 2, 3, 4}
	tes = append(tes, 5, 6)

	out := []int{12, 13}

	fmt.Println(tes)

	tes = append(tes, out[1:]...)

	fmt.Println(tes)

	tes2 := [][]int{
		{7, 8, 9},
		{10, 11},
		tes,
	}

	fmt.Println(tes2)
}
