/*
CANAL BIDIRECCIONAL
chan type
	Puede enviar y recibir informacion
*/

package main

import "fmt"

func main() {
	ch := make(chan bool) // creamos un canal
	go myRoutine(ch) // mandamos llamar la routine y le pasamos el canal
	result := <- ch // esperamos a que el canal nos envie un valor
	fmt.Println(result) // imprimimos el valor que envia el canal
}

func myRoutine(ch chan bool) {
	ch <- true // hacemos que el canal reciba un valor
}
