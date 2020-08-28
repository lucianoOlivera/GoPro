package main

import (
	"fmt"
)

func main() {
	par := make(chan int)
	impar := make(chan int)
	salir := make(chan bool)

	go enviar(par, impar, salir)

	recibir(par, impar, salir)

	fmt.Println("Finalizar.")
}

// send channel (canal para enviar)
func enviar(par, impar chan<- int, salir chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
	close(salir)
}

// receive channel (canal para recibir)
func recibir(par, impar <-chan int, salir <-chan bool) {
	for {
		select {
		case v := <-par:
			fmt.Println("El valor recibido del canal par:", v)
		case v := <-impar:
			fmt.Println("El valor recibido del canal impar:", v)
		case i, ok := <-salir:
			if !ok {
				fmt.Println("desde comma ok", i, ok)
				return
			} else {
				fmt.Println("desde comma ok", i)
			}
		}
	}
}
