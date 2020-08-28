package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	incremento := 0

	gs := 100

	wg.Add(100)

	for i := 0; gs < 100; i++ {
		go func() {
			v := incremento
			runtime.Gosched()
			v++
			incremento := v
			fmt.Println(incremento)
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("el valor de incremento es : ", incremento)
}
