/* 
La declaración select extrae el valor de cualquier canal que tenga un valor listo para ser
extraído.
*/

package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	exit := make(chan int)

	go send(even, odd, exit)
	receive(even, odd, exit)

	fmt.Println("Finishing")
}

func send(even, odd, exit chan<- int) {
	for j := 0; j < 100; j++ {
		if j % 2 == 0 {
			even <- j
		} else {
			odd <- j
		}
	}

	close(even)
	close(odd)
    
	exit <- 0
}

func receive(even, odd, exit <-chan int) {
	for {
		/* Select es como si tuvieramos un switch donde podemos verificar
		si tenemos valores en los canales */
		select {
			case i := <-even: // si tenemos valores en el canal par
                fmt.Println("Desde el canal par", i)
            case i := <-odd: // si tenemos valores en el canal impar
                fmt.Println("Desde el canal impar", i)
			case i:= <-exit: // si tenemos valor en el canal salir
                fmt.Println("Salimos del programa", i)
				return
		}
	}
}
