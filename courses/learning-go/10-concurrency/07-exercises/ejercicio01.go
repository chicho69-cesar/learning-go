/*
Además de la gorutina principal, lanza dos gorutinas adicionales
	Cada gorutina debe imprimir algo en pantalla
Usa waitgroups para asegurarte que cada gorutina finalize antes de que el programa
termine
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Numero de CPUs: ", runtime.NumCPU())
	fmt.Println("Numero de GoRoutines: ", runtime.NumGoroutine())

	wg.Add(2)

	go func() {
		fmt.Println("Hola desde la primera goroutine")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hola desde la segunda goroutine")
		wg.Done()
	}()

	fmt.Println("Numero de CPUs: ", runtime.NumCPU())
	fmt.Println("Numero de GoRoutines: ", runtime.NumGoroutine())

	fmt.Println("\nA punto de finalizar")
	wg.Wait()
}
