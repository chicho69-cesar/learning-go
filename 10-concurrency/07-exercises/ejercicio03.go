/*
Usando gorutinas, crea un programa incremento
	Haz que una variable guarde el valor del incremento
	Lanza varias gorutinas
		cada gorutina deberá
			Leer el valor del incremento
				Almacenarlo en una nueva variable
			Ceder el procesador con runtime.Gosched()
			Incrementa la nueva variable
			Escribe el valor en la nueva variable de vuelta a la variable
			incremento
Usa waitgroups para esperar que todas las gorutinas finalicen
Lo anterior generará una race condition.
Comprueba que es una race condition usando el flag -race
Si necesitas ayuda, aquí tienes una pista: https://play.golang.org/p/a-tdD-7lTId
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	incrementer := 0
	limit := 100
	wg.Add(limit)

	for i := 0; i < limit; i++ {
		go func() {
			value := incrementer
			value++
			incrementer = value
			runtime.Gosched()

			fmt.Println("Incremento: ", incrementer)

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("El valor final es: ", incrementer)
}
