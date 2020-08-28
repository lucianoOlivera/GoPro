package main

import "fmt"

/*Escribe un programa que:
Lance 10 gorutinas
Cada gorutina agrega 10 números a un canal
Extrae los números del canal e imprímelos*/

func main() {
	c := make(chan int)

	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
	}

	for k := 0; k < 100; k++ {
		fmt.Println(k, <-c)
	}

	fmt.Println("A punto de finalizar.")
}
