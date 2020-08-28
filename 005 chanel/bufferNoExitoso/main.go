package main

import "fmt"

func main() {
	ca := make(chan int, 1)

	ca <- 64
	ca <- 65

	fmt.Println(<-ca)

}
