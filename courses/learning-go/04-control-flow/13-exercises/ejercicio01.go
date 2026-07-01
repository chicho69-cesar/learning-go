/* 
Imprime todos los números del 1 al 10,000
*/

package main

import "fmt"

func main() {
	for i := 1; i <= 10000; i++ {
		fmt.Printf("%d\t", i)

		if i % 100 == 0 {
			fmt.Println()
		}
	}
}
