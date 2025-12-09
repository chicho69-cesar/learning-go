/* 
Usando iota, crea 4 constantes para los PRÓXIMOS 4 años. Imprime 
los valores de las constantes.
*/

package main 

import "fmt"

const (
	_ = iota
	year1 = iota + 2022
	year2 = iota + 2022
	year3 = iota + 2022
	year4 = iota + 2022
)

func main() {
	fmt.Printf("Año 1: %d\n", year1)
	fmt.Printf("Año 2: %d\n", year2)
	fmt.Printf("Año 3: %d\n", year3)
	fmt.Printf("Año 4: %d\n", year4)
}
