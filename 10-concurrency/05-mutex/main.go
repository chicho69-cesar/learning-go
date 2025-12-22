/* 
Un "mutex" es un bloqueo de exclusión mutua. Los mutexes nos permiten bloquear nuestro
código para que solo una gorutina pueda acceder a ese bloque de código bloqueado a la
vez.
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Número de CPUs:", runtime.NumCPU())
	fmt.Println("Número de GoRoutines:", runtime.NumGoroutine())

	counter := 0
	const countLimit = 100
	var wg sync.WaitGroup

	wg.Add(countLimit)

	var mut sync.Mutex

	for i := 0; i < countLimit; i++ {
		go func() {
			mut.Lock() // Bloqueamos las goroutines hasta que esta se desbloquee

			value := counter
			value++
			runtime.Gosched() 
			counter = value

			mut.Unlock() // desbloqueamos la goroutine

			wg.Done()
		}()

		fmt.Println("Número de GoRoutines:", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Número de GoRoutines:", runtime.NumGoroutine())
	fmt.Println("Contador: ", counter)
}
