package testes

import "fmt"

func MaisFunc() {

	mak := make([]int, 10, 20)

	fmt.Println(mak)

}

func Maps() {

	teste := map[string]int{
		"Paulo": 12345,
		"outro": 67890,
	}
	teste["Mais um"] = 98765

	fmt.Println(teste)
	fmt.Println(teste["Paulo"])

	//comma ok idiom
	if a, ok := teste["Rodrigo"]; !ok {
		fmt.Println("deu erro")
	} else {
		fmt.Println(a)
	}
	delete(teste, "outro")
	//range
	for key, value := range teste {
		fmt.Println(key, value)
	}
}
