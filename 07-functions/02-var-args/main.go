/*
Puedes crear una func el cual toma un número ilimitado de argumentos. Cuando haces
esto, se conoce como “parámetro variado” o en inglés como “variadic parameter.” Cuando usas
el operador que forma parte de los elementos de léxico “...T” para indicar un parámetro variado
(donde “T” representa algún tipo).
*/

package main

import "fmt"

func main() {
	nums := []int { 1, 2, 3, 4, 5 }
	suma1 := sum(nums...) // podemos pasar un slice 
	fmt.Printf("\n\nLA SUMA 1 ES: %d\n", suma1)

	suma2 := sum(6, 7, 8, 9, 0) // podemos pasar todos los valores que queramos
	fmt.Printf("\n\nLA SUMA 2 ES: %d\n", suma2)
}

// recibimos una cantidad variada de parametros en esta función
func sum(nums ...int) int {
	fmt.Println(nums)
	fmt.Printf("%T: \n", nums)

	suma := 0

	for i, v := range nums {
		suma += v
		fmt.Println("El valor en el índice", i, "suma", v, "al total, quedando", suma)
	}

	fmt.Println("\nEl valor de la suma es: ", suma)

	return suma
}
