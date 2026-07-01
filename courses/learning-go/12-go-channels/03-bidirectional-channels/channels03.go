/*
CANALES SEND ONLY
<-chan int
	canales en los que solo se puede enviar valores, es decir, este canal
	solo nos sirve para almacenar valores en el
*/

package main

// import "fmt"

func main() {
	ch := make(chan<- int)

	go func() {
		ch <- 21 // no se le puede enviar informacion a un canal receive only
	}()

	// fmt.Println(<-ch) // no podemos recibir informacion de un canal que
	// es solo para enviar informacion
}
