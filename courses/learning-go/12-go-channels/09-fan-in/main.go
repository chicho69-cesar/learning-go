/* 
Tomando valores de varios canales y colocándolos en un canal.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send(even, odd)
	go receive(even, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("Finishing...")
}

func send(even, odd chan<- int) {
	for i := 0; i < 100; i++ {
		if i % 2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}

	close(even)
    close(odd)
}

func receive(even, odd <-chan int, fanin chan<- int) {
    var wg sync.WaitGroup
	wg.Add(2)

	go func() {
        for i := range even {
			fanin <- i
		}

		wg.Done()
	}()

	go func() {
        for i := range odd {
			fanin <- i
		}
		
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}
