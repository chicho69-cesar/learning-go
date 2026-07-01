/* 
Arregla la race condition que creaste en el ejercicio anterior usando un mutex
	Tiene sentido eliminar runtime.Gosched()
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	incrementer := 0
	limit := 100
	wg.Add(limit)

	var mut sync.Mutex

	for i := 0; i < limit; i++ {
		go func() {
			mut.Lock()
			
			value := incrementer
			value++
			incrementer = value
			runtime.Gosched()

			fmt.Println("Incremento: ", incrementer)

			mut.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("El valor final es: ", incrementer)
}
