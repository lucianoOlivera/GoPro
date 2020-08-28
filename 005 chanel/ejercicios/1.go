package main

/*"Ejercicio Práctico #1
Haz que este código funcione:
Usando una función literal, también conocida como, función anónima autoejecutable
Solución: https://play.golang.org/p/SHr3lpX4so 
Un canal con búfer
Solución: https://play.golang.org/p/Y0Hx6IZc3U */



import (
	"fmt"
)

func main() {
	c := make(chan int,1)

	c <- 42

	fmt.Println(<-c)
}
