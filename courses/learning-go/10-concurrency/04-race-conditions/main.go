package main

import (
	"fmt"
	"runtime"
	"sync"
	// "time"
)

func main() {
	fmt.Println("Número de CPUs:", runtime.NumCPU())
	fmt.Println("Número de GoRoutines:", runtime.NumGoroutine())

	counter := 0
	const countLimit = 100
	var wg sync.WaitGroup

	wg.Add(countLimit)

	for i := 0; i < countLimit; i++ {
		go func() {
			value := counter
			value++

			// time.Sleep(time.Second * 2)
			runtime.Gosched() // esto es como hacer yield en otros lenguajes

			counter = value
			wg.Done()
		}()

		fmt.Println("Número de GoRoutines:", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Contador: ", counter)
}
