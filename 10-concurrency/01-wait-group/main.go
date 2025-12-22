/*
Un WaitGroup espera que una colección de gorutinas terminen su trabajo. La gorutina de
main llama Add para configurar el número de gorutinas por las que tiene que esperar. Luego
cada una de las gorutinas corre y llama a Done cuando terminan. Al mismo tiempo, Wait puede
ser usado para bloquear hasta que todas las gorutinas hayan finalizado. Escribir código
concurrente es súper fácil: todo lo que tenemos que hacer es poner “go” en frente de una
llamada a una función o método.
	runtime.NumCPU()
	runtime.NumGoroutine()
	sync.WaitGroup
		func (wg *WaitGroup) Add(delta int)
		func (wg *WaitGroup) Done()
		func (wg *WaitGroup) Wait()
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup // creamos nuestro wait group

func main() {
	fmt.Println("Operating System: \t\t", runtime.GOOS) // Sistema Operativo
	fmt.Println("Architecture: \t\t", runtime.GOARCH) // Arquitectura del sistema
	fmt.Println("Number of CPUs: \t\t", runtime.NumCPU()) // Numero de cores
	fmt.Println("GoRoutines: \t\t", runtime.NumGoroutine()) // Numero de rutinas

	wg.Add(1) // le decimos a nuestro wait group que vamos a tener una rutina

	go foo() // lanzamos la rutina
	bar()

	fmt.Println("Number of CPUs: \t\t", runtime.NumCPU()) 
	fmt.Println("GoRoutines: \t\t", runtime.NumGoroutine())

	wg.Wait() // le decimos a main que espere a que terminen todas las rutinas
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}

	wg.Done() // informamos de que la rutina termino
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
