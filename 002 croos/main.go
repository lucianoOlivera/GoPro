package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("sistema operativo : %v \narquitectura: %v\n", runtime.GOOS, runtime.GOARCH)

}
