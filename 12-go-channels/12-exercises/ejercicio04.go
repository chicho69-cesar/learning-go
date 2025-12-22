/* 
Comenzando con este código, extrae los valores del canal usando una declaración
select
*/

package main

import (
	"fmt"
)

func main() {
	salir := make(chan int)
	c := generate(salir)

	receive(c, salir)

	fmt.Println("A punto de finalizar.")
}

func generate(salir chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}

		close(salir)
	}()

	return c
}

func receive(c <-chan int, salir <-chan int) {
	for {
		select {
			case v := <-c:
				fmt.Println(v)
			case v, ok := <-salir:
				if !ok {
					fmt.Println("Salimos", v)
                    return
                }

                fmt.Println("Continuamos", v)
        }
	}
}

/* Otra alternativa */
/* package main

import (
	"fmt"
)

func main() {
	salir := make(chan int)
	c := gen(salir)

	recibir(c, salir)

	fmt.Println("A punto de finalizar.")
}

func gen(salir chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		salir <- 1
		close(c)
	}()

	return c
}

func recibir(c, salir <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println("Valor: ", v)
		case <-salir:
			return
		}
	}
}
 */
