package main

import (
	"fmt"
)

func main() {
	x := 1

	for {
		x++
		if x > 100 {
			break // break rompe con el ciclo
		}

		if x % 2 != 0 {
			continue // saltamos a la siguiente ejecucion
		}

		fmt.Println(x)
		
	}
	
	fmt.Println("done.")
}
