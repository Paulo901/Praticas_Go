package exercicios

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("Sistema Operacional: ", runtime.GOOS)
	fmt.Println("Arquivo: ", runtime.GOARCH)
}
