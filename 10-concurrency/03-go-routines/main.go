/*
Son llamadas goroutinas porque los términos existentes —threads, coroutines, procesos, entre
otros — transmiten connotaciones inexactas. Una goroutine tiene un modelo simple: es una
función ejecutándose concurrentemente con otras goroutinas en el mismo espacio de memoria
(address space).
*/

package main

import "fmt"

func main() {
	ch := make(chan int) // creamos un canal que va a recibir numeros int
	
	/* Empleamos la palabra go para mandar una función a ejecutarse en
	una goroutine */
	go func() {
		ch <- doSomething(5) // le asignamos el resultado que esperamos de la función al canal
	}()

	fmt.Println(<-ch) /* el operador <- es igual al await ya que con
	el esperamos a que la goroutine termine */
}

func doSomething(x int) int {
	return x * 5
}
