package exercicios

import (
	"fmt"
	"runtime"
)

func Exer6() {

	fmt.Println("Sistema Operacional: ", runtime.GOOS)
	fmt.Println("Arquivo: ", runtime.GOARCH)
}
