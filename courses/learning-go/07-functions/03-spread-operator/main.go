/* 
Cuando tienes un slice de algún tipo, puedes pasar cada uno de los valores individuales en
un slice usando el operador “...”
*/

package main

import "fmt"

func main() {
	x := sum("james", 1, 2, 3, 4)
	fmt.Println("The total is", x)
}

/* Cuando usamos var args en una función el parámetro que es el var arg
debe estar especificado hasta el final de la lista de parametros */
func sum(s string, x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	sum := 0

	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total which is now", sum)
	}

	fmt.Println("The total is", sum)

	return sum
}
