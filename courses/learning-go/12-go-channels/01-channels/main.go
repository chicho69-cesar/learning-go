package main

import "fmt"

func main() {
	// creamos un canal con un buffer de 2, para tener 2 goroutines 
	// donde una envia un valor por el canal hacia la otra
	/* el buffer es la cantidad de elementos que se pueden quedar
	dentro del canal */
	ch := make(chan int, 2)
	ch <- 20

	fmt.Println(<-ch)

	ch <- 21
	ch <- 22 // tenemos dos goroutines en el canal

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
