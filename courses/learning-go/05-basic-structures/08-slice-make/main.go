/* 
Los slices son construidos sobre los arreglos. Un slice es dinámico, así este crecerá en tamaño.
Sin embargo, el arreglo subyacente, no crece en tamaño. Cuando creamos un slice, podemos
usar la función predefinida interna make para especificar que tan grande debería ser
nuestro slice y también qué tan grande debería ser el arreglo subyacente. Esto puede
mejorar un poco el desempeño del programa.
	make([]T, length, capacity)
	make([]int, 50, 100)
*/

/* 
Cuando en un slice el tamaño supera a la capacidad, lo que Go hace es que 
duplica la capacidad del slice que fue definida, si la capacidad no fue 
definida esta se ira tomando como potencias de 2
Capacidades: 
2 4 8 16 32 64 128 256 ...
*/

package main

import "fmt"

func main() {
	x := make([]int, 10, 12) // slice int con tamaño 10 y capacidad 12
	fmt.Println(x)
	fmt.Println(len(x)) // 10
	fmt.Println(cap(x)) // 12
	x[0] = 42
	x[9] = 999

	fmt.Println(x)
	fmt.Println(len(x)) // 10
	fmt.Println(cap(x)) // 12

	x = append(x, 3423) // agregamos al final

	fmt.Println(x)
	fmt.Println(len(x)) // 11
	fmt.Println(cap(x)) // 12

	x = append(x, 3424) // agregamos al final

	fmt.Println(x)
	fmt.Println(len(x)) // 12 
	fmt.Println(cap(x)) // 12

	x = append(x, 3425) // agregamos al final

	fmt.Println(x)
	fmt.Println(len(x)) // 13
	fmt.Println(cap(x)) // 24
}
