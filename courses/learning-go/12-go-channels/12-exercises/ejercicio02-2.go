/* 
Haz que este código funcione:
# https://play.golang.org/p/oB-p3KMiH6
	Solución: https://play.golang.org/p/isnJ8hMMKg
# https://play.golang.org/p/_DBRueImEq
	Solución: https://play.golang.org/p/mgw750EPp4
*/

package main

import (
	"fmt"
)

func main() {
	cr := make(chan int)

	go func() {
		cr <- 42
	}()
	
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}
