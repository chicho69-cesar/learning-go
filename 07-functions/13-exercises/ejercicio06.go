/* 
Crea y usa una func anónima que imprima los numeros del 1 al 100
*/

package main

import "fmt"

func main() {
	func() {
		for i := 1; i <= 100; i++ {
			fmt.Println(i)
		}
	}()

	fmt.Println("\nListo!")
}
