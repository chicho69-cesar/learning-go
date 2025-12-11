//* En Go no existen ni while ni do-while, solamente for

package main

import (
	"fmt"
)

func main() {
	// Estructura de un cliclo 
	/* for init; condition; post {
		fmt.Println("Hello, playground")
	} */

	// for init; condition; post {}
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
}
