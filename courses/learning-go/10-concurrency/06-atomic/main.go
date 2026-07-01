/* 
Podemos usar el paquete atomic para también prevenir una race condición en nuestro
código incremental.
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("Número de CPUs:", runtime.NumCPU())
	fmt.Println("Número de GoRoutines:", runtime.NumGoroutine())

	var counter int64
	const countLimit = 100
	var wg sync.WaitGroup

	wg.Add(countLimit)

	for i := 0; i < countLimit; i++ {
		go func() {
			// agregamos un valor a la direccion del counter
			atomic.AddInt64(&counter, 1)
			runtime.Gosched() 
			// cargamos el valor de la direccion del counter
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			wg.Done()
		}()

		fmt.Println("Número de GoRoutines:", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Número de GoRoutines:", runtime.NumGoroutine())
	fmt.Println("Contador: ", counter)
}
