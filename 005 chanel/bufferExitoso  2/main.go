package main

import "fmt"

func main() {
	ca := make(chan int, 2)

	ca <- 64
	ca <- 54
	fmt.Println(<-ca)
	fmt.Println(<-ca)
}
