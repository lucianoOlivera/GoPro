package main

import "fmt"

func main() {
	par := make(chan int)
	impar := make(chan int)
	salir := make(chan int)

	go enviar(par, impar, salir)

	recibir(par, impar, salir)

}

func enviar(p, i, s chan<- int) {
	for j := 0; j < 100; j++ {
		if j%2 == 0 {
			p <- j
		} else {
			i <- j
		}
	}

	s <- 0

}

func recibir(p, i, s <-chan int) {

	for {
		select {
		case v := <-p:
			fmt.Println("desde el canal par", v)
		case v := <-i:
			fmt.Println("desde el canal impar", v)
		case v := <-s:
			fmt.Println("desde el canal salir", v)
			return
		}

	}
}
