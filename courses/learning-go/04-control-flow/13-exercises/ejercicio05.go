/* 
Imprime el resto o módulo, el cual es resultado de dividir entre 4 cada 
número en el rango de 10 y 100.
*/

package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("El residuo de %d / 4 = %d\n", i, i % 4)
	}
}
