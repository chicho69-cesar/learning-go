/*
Crea una func el cual retorna una func
Asigna la func retornada a una variable
Llama la func retornada
*/

package main

import "fmt"

func main() {
	counter1 := clousure(5)
	counter2 := clousure(10)

	fmt.Println(counter1())
	fmt.Println(counter1())
	fmt.Println(counter2())
	fmt.Println(counter2())
	fmt.Println(counter1())
	fmt.Println(counter1())
	fmt.Println(counter2())
}

func clousure(init int) func() int {
	state := init

	return func() int {
		state++
		return state
	}
}
