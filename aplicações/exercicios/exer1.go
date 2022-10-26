package exercicios

import (
	"encoding/json"
	"fmt"
)

// código do link: https://play.golang.org/p/U0jea43X55
type user struct {
	First string
	Age   int
}

func Exer1() {
	U1 := user{
		First: "James",
		Age:   32,
	}

	U2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	U3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{U1, U2, U3}

	fmt.Println(users)

	// fim do código do link

	emjson, err := json.Marshal(users)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(emjson))
}
