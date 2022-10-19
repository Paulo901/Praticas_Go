package exercicios

import "fmt"

var y = []int{40, 41, 42, 43}

func Exer4() {

	y = append(y, 44)
	y = append(y, 45, 46, 47)

	fmt.Println(y)

	x := []int{48, 49, 50}

	y = append(y, x...)

	fmt.Println(y)
}
