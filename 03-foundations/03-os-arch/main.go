package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Accedemos al sistema operativo donde corre el programa
	fmt.Println(runtime.GOOS)
	// Accedemos a la arquitectura de este sistema operativo
	fmt.Println(runtime.GOARCH)
}
