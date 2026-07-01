/*
Escribe un programa que:
	Ponga 100 números en un canal
	Extraiga los números del canal y los imprima
*/

package main

import "fmt"

func main() {
	ch := make(chan int)
	go add(ch)
	print(ch)

	fmt.Println("Finalizado")
}

func add(ch chan<- int) {
	for i := 1; i <= 100; i++ {
		ch <- i
	}

	close(ch)
}

func print(ch <-chan int) {
	for v := range ch {
		fmt.Println("Valor: ", v)
	}
}
