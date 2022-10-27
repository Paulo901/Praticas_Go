package exercicios

import (
	"encoding/json"
	"fmt"
	"os"
)

// código referente a link	-> https://go.dev/play/p/BVRZTdlUZ_ <-
func Exer3() {
	//	usando recursos de noRep.go

	u1 := User{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := User{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := User{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []User{u1, u2, u3}

	//	fim do código do link
	//---------------------------

	enc := json.NewEncoder(os.Stdout) //.Encode(users)
	err := enc.Encode(users)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(enc)

}
