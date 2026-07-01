package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	exit := make(chan bool)

	go send(even, odd, exit)
	recieve(even, odd, exit)

	fmt.Println("Finishing.")
}

func send(even, odd chan<- int, exit chan<- bool) {
	for i := 0; i < 100; i++ {
		if i % 2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}

    close(exit) /* Cuando cerramos el canal exit, lo que hacemos es que
	al final de la goroutine cerramos el canal exit y esto le agrega un exit
	con un cero por default */
}

func recieve(even, odd <-chan int, exit <-chan bool) {
	for {
		select {
			case i := <-even:
				fmt.Println("Even: ", i)
			case i := <-odd:
				fmt.Println("Odd: ", i)
			case i, ok := <-exit: // obtenemos el valor del canal y si esta cerrado
				if !ok { // si esta cerrado marca false
                    fmt.Println("Desde coma ok - if", i)
					return
                } else {
                    fmt.Println("Desde coma ok - else", i)
                }
        }
	}
}
