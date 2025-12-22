/*
Consiste en tomar algún trabajo a realizar y convertirlo en varias porciones de trabajo con
varias gorutinas.
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go add(c1)
	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("About to exit...")
}

func add(c chan int) {
	for i := 0; i < 100; i++ {
        c <- i
    }
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup

	for v := range c1 {
		wg.Add(1)
		go func(value int) {
			c2 <- workThatConsumeTime(value)
			wg.Done()
		}(v)
	}

	wg.Wait()
	close(c2)
}

func workThatConsumeTime(num int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return num * rand.Intn(1000)
}
