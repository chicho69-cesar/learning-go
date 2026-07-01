/*
Arreglar la race condition que creaste en el ejercicio #4 usando el paquete atomic
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup

	var incrementer int64
	limit := 100
	wg.Add(limit)

	for i := 0; i < limit; i++ {
		go func() {
			atomic.AddInt64(&incrementer, 1)
			runtime.Gosched()

			fmt.Println("Incremento: ", atomic.LoadInt64(&incrementer))

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("El valor final es: ", incrementer)
}
