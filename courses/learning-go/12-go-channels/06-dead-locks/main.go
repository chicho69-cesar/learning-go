package main

import "fmt"

func main() {
	ch := make(chan int)

	go send(ch)
	receive(ch)

	fmt.Println("Finishing")
}

func send(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}

	close(ch)
}

func receive(ch <-chan int) {
    for v := range ch {
		fmt.Println(v)
	}
}
