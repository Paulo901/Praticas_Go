package exercicios

import "fmt"

func Exer6() {
	func() {
		fmt.Println(2)
	}()
}
func Exer7() {
	t := func() int {
		return 1
	}
	fmt.Println(t())
}
