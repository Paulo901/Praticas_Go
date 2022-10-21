package testes

// retornando uma função
func Ola(nome string) func(t string) (string, string) {

	return func(turno string) (string, string) {

		switch turno {
		case "manha":
			return "Bom dia, ", nome
		case "tarde":
			return "Boa tarde, ", nome
		case "noite":
			return "Está tarde! ", nome
		default:
			return "turno inexistente, ", nome
		}
	}
}

// pegando função como parâmetro
func Soma1(nums ...int) int {
	var result int = 0
	for _, valor := range nums {
		result += valor
	}
	return result
}

func SoImpar(f func(nums ...int) int, onums ...int) int {
	impares := []int{}
	for _, v := range onums {
		if v%2 != 0 {
			impares = append(impares, v)
		}
	}
	res := f(impares...)
	return res
}
