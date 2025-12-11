/* 
Aquí mostramos como agregar un elemento a un mapa. También muestro cómo usar el ciclo
for range para imprimir las llaves y valores de un mapa.
*/

package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}

	fmt.Println(m)

	fmt.Println(m["James"])
	fmt.Println(m["Barnabas"])

	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	m["todd"] = 33 // Agregamos valor al map

	if v, ok := m["Barnabas"]; ok {
		fmt.Println(v)
	}

	/* El range dentro de un map, al igual que en slices nos ayuda
	a iterar la estructura de datos, en los slice devuelve index y value
	aqui devuelve key y value */
	for k, v := range m {
		fmt.Println(k, v) // key - value
	}

	xi := []int{ 4, 5, 7, 8, 9, 42 }

	for i, v := range xi {
		fmt.Println(i, v) // index - value
	}
}
