/* 
recover() es una función interna que toma el control de una goroutine
que recupera el control de una goroutine en panico, recover solo es 
util dentro de una función diferida, durante una función normal
el llamado a recover solo devolvera nil y no tendra ningun otro efecto
*/

package main

import (
	"fmt"
)

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		// nos recuperamos con recover
		if err := recover(); err != nil {
            fmt.Println("Recovered from panic:", err)
        }
	}()

	fmt.Println("Calling g.")
    g(0)
	fmt.Println("Returned normally from g.")
}

func g(num int) {
    if num > 3 {
		fmt.Println("Panicking...!")
		panic(fmt.Sprintf("%v", num))
	}

	defer fmt.Println("Defer in g", num)
	fmt.Println("Printing in g", num)

	g(num + 1)
}
