/* 
Podemos eliminar elementos de un slice usando append y slicing (dividiendo). Este es un
maravilloso y elegante ejemplo de porqué Go es súper cool y cómo provee facilidad de
programación.
*/

package main

import "fmt"

func main() {
	x := []int{ 4, 5, 7, 8, 42 }
	fmt.Println(x)
	x = append(x, 44, 55, 66)
	fmt.Println(x)

	/* x hasta aqui -> 4, 5, 7, 8, 42, 44, 55, 66 */

	y := []int{ 333, 444, 666 }
	x = append(x, y...)
	fmt.Println(x)

	/* x hasta aqui -> 4, 5, 7, 8, 42, 44, 55, 66, 333, 444, 666 */

	/* Aqui lo que hacemos es que al slice que se forma con elementos de
	la posicion 0 hasta antes de la 2 (4, 5) se le agregan los elementos 
	del slice de la posicion 4 hasta el final (42, 44, 55, 66 ...) */
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

	/* x hasta aqui -> 4, 5, 42, 44, 55, 66, 333, 444, 666 */
}
