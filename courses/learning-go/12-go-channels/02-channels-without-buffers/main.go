package main

import "fmt"

func main() {
	/* creamos un canal sin buffer para almacenar valores enteros */
	ch := make(chan int)

	/* mandamos a llamar a una goroutine, para mandar un valor hacia 
	el canal con un valor entero */
	go func() {
		ch <- 21
	}()

	/* almacenamos el valor que tenemos en un canal */
	result := <- ch

	fmt.Println(result)
}
