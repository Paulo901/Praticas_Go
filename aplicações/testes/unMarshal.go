package testes

import (
	"encoding/json"
	"fmt"
)

type coisa2 struct {
	Algo  string `json: "algo"`
	Nivel int    `json: "nivel"`
}

func TirandoDoMarshal() {

	sb := []byte(`{"Algo":"monstro","Nivel":100}`)

	var c2 coisa2

	err := json.Unmarshal(sb, &c2)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(c2)

	//	com UnMarshal
	//-------------------------------------
	//	com Decoder

}
