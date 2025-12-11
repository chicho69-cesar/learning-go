/* 
Usando el código del ejercicio anterior, usa SLICING para crear los siguientes nuevos slices el
cual serán impresos:
	[42 43 44 45 46]
	[47 48 49 50 51]
	[44 45 46 47 48]
	[43 44 45 46 47]
*/

package main

import "fmt"

func main() {
	slice := []int{ 42, 43, 44, 45, 46, 47, 48, 49, 50, 51 }

	sl1 := slice[:5]
	fmt.Println(sl1)

	sl2 := slice[5:]
	fmt.Println(sl2)

	sl3 := slice[2:7]
	fmt.Println(sl3)

	sl4 := slice[1:6]
	fmt.Println(sl4)
}
