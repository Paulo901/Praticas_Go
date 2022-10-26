package testes

import (
	"encoding/json"
	"fmt"
	"os"
)

type coisa struct {
	Algo  string
	Nivel int
}

func Transforma() {

	c1 := coisa{Algo: "monstro", Nivel: 100}

	c1j, err := json.Marshal(c1)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(c1j))
	//	com Marshal
	//-------------------------------------
	//	com Encoder
	encodador := json.NewEncoder(os.Stdout)
	encodador.Encode(c1)

}
