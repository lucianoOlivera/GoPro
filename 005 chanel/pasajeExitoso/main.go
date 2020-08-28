package main

import "fmt"

func main() {
	ca := make(chan int)
	go func() {
		ca <- 64
	}()
	fmt.Println(<-ca)

}
