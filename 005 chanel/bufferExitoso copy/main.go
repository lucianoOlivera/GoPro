package main

import "fmt"

func main() {
	ca := make(chan int, 1)

	ca <- 64

	fmt.Println(<-ca)

}
