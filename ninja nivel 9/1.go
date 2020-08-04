package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() { fmt.Println("Desde la rutina 1") wg.Done()
	}()
	go func() { fmt.Println("Desde la rutina 2") wg.Done()
	}()
	fmt.Printf("Desde el main")
	wg.Wait()
}
