package testes

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var Wg sync.WaitGroup

func Prin() {

	Wg.Add(1)

	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	Primeira()

	for i := 1; i <= 20; i++ {
		fmt.Println("Func pri")
	}
	fmt.Println(runtime.NumGoroutine())
}
func Primeira() {

	for i := 1; i <= 20; i++ {
		fmt.Println("Func 1")
		time.Sleep(time.Duration(20))
	}

}
