/*
CANALES RECIEVE ONLY
<-chan int
	canales solo para recibir valores, es decir, este canal
	solo nos sirve para tomar valores almacenados en el
*/

package main

import "fmt"

func main() {
	ch := make(<-chan int)

	go func() {
		// ch <- 21 // no se le puede enviar informacion a un canal receive only
	}()

	fmt.Println(<-ch)
}
