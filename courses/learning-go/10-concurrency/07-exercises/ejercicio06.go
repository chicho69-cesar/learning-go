/*
Crea un programa que imprima tu OS y ARCH.
*/

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Sistema Operativo: ", runtime.GOOS)
	fmt.Println("Arquitectura: ", runtime.GOARCH)
}
