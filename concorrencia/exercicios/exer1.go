package exercicios

import (
	"fmt"
	"sync"
	"time"
)

func Exer1() {

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		for i := 1; i <= 20; i++ {
			fmt.Println("primeira")
			time.Sleep(time.Duration(20))
		}
	}()

	go func() {
		for i := 1; i <= 20; i++ {
			fmt.Println("segunda")
			time.Sleep(time.Duration(20))
		}
		wg.Done()
	}()

	wg.Wait()
}
