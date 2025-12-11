/* 
Podemos iterar sobre los valores en un slice con la cláusula range. También le podemos
agregar elementos a un slice mediante el uso del índice.
*/

package main

import "fmt"

func main() {
	x := []int{ 1, 2, 3, 4, 5 }
	fmt.Println(x)
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])

	for index, value := range x {
		fmt.Println(index, " ", value)
	}

	fmt.Println()

	for _, value := range x {
		fmt.Println(value)
	}
}
