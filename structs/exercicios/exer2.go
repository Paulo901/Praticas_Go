package exercicios

import "fmt"

func Exer2() {

	p1 := map[string]pessoa{
		"Oliveira": {nome: "Paulo", sobrenome: "Oliveira", hobby: []string{"Xadrez", "cubo magico"}},
		"Fulano":   {nome: "Outro", sobrenome: "Fulano", hobby: []string{"natação"}},
	}

	for key, value := range p1 {
		fmt.Println("O ", value.nome, " ", key, "Gosta de: ")
		for _, valor := range value.hobby {
			fmt.Println(" -> ", valor)
		}

	}

}
