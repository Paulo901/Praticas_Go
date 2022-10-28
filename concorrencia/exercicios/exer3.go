package exercicios

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var contador int

func Exer3() {

	A(3)
	wg.Wait()

	fmt.Println(contador)

}

func A(n int) {
	for i := 0; i < n; i++ {
		go func() {
			c := contador
			runtime.Gosched()
			c++
			contador = c
		}()
	}
}
