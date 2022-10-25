package exercicios

func Exer10() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
