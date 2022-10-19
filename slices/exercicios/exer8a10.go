package exercicios

import "fmt"

func Exer8() {

	b := map[string]string{
		"Paulo Rodrigo": "Xadrez",
		"Outro Fulano":  "Nenhum",
	}

	for key, value := range b {
		fmt.Println(key, " Hobby: ", value)
	}
	b["Alguem"] = "Fazer nada"

	fmt.Println(b)

	delete(b, "Alguem")

	fmt.Println(b)
}
