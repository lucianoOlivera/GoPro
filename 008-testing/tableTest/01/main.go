package main

import "fmt"

func main() {
	fmt.Print("la suma de dos mas 4 : ", Mysuma(2, 4))
	fmt.Print("la suma de 1 mas 5 :", Mysuma(1, 5))
	fmt.Print("la suma de 2 mas 2 :", Mysuma(2, 2))

}
func Mysuma(xi ...int) int {
	var sum int
	for _, v := range xi {
		sum += v
	}
	return sum
}
