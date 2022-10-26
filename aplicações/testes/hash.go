package testes

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Criptografia() {

	senha := "Paulo123"
	Esenha := "paulo123"

	sb, err := bcrypt.GenerateFromPassword([]byte(senha), 10)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(Esenha)
	fmt.Println(sb)
	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senha)))
	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(Esenha)))

}
