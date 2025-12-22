/* 
Parte del paquete testing nos permite medir la velocidad de nuestro código. Esto también
podría ser llamado “midiendo el rendimiento” de tu código, o “benchmarking” tu código -
averiguar qué tan rápido se ejecuta el código.

para correr los tests que tienen pruebas de rendimiento usamos el comando: 
	go test -bench .
*/

package main

import "fmt"

func main() {
	fmt.Println(Saludar("Cesar"))
}

func Saludar(s string) string {
	return fmt.Sprint("Bienvenido querido ", s)
}
