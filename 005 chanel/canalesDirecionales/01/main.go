package main

import "fmt"

func main() {
	ca := make(chan int, 2)

	ca <- 64
	ca <- 65

	fmt.Println(<-ca)
	fmt.Println("----------------")
	fmt.Println("%T\n", ca)

}
