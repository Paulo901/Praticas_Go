package exercicios

import "fmt"

func Exer4() {

	type a struct {
		mapa  map[int]string
		slice []int
	}

	algo := a{mapa: map[int]string{1: "algo"}, slice: []int{1, 2, 3, 4}}
	fmt.Println("mostrando: ", algo)

}
