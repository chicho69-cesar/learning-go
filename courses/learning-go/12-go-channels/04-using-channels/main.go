/*
Crear un canal general
En funciones puedes especificar
	Canal receive-only
		Puedes recibir valores del canal
		Un parámetro canal receptor
		En la función, puedes solamente extraer valores del canal
		No puedes cerrar un canal receptor
	Canal send-only
		Puedes enviar valores al canal
		No puedes recibir/extraer/leer desde el canal
		Solamente puedes enviar valores al can
*/

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	ch := make(chan int)
	go send(ch) // lanzamos la goroutine send
	go receive(ch) // lanzamos la goroutine recieve

	wg.Wait()

	fmt.Println("\nPROGRAMA FINALIZADO")
}

// recibimos un canal de send only
func send(ch chan<- int) {
	ch <- 21
	wg.Done()
}

// recibimos un canal de recieve only
func receive(ch <-chan int) {
	var value int = <-ch
	fmt.Println("El valor almacenado en el canal es: ", value)
	wg.Done()
}
