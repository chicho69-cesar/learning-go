/* 
Crea un programa que muestre el “if statement” en acción.
*/

package main

import "fmt"

func main() {
	var age int

	fmt.Print("Escribe tu edad: ")
	fmt.Scan(&age)

	if (age < 18) {
		fmt.Println("No eres mayor de edad :(")
	}
}
