package exercicios

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var mutex sync.Mutex
var contador int32

func Exer3a5() {

	A(3)
	wg.Wait()

	fmt.Println(contador)

}

func A(n int) {
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			//mutex.Lock() 	//(feito no exer4)
			atomic.AddInt32(&contador, 1)
			//  c := contador 	//(feito no exer3)
			//	runtime.Gosched() // (a localização muda com o atomic)
			//c++ 	//(feito no exer3)
			atomic.LoadInt32(&contador)
			//mutex.Unlock() //(feito no exer4)
			runtime.Gosched()
			wg.Done()
		}()
	}
}
