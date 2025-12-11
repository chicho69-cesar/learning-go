/* 
Usando el ejercicio anterior, crea un programa que use “else if” y “else”.
*/

package main

import "fmt"

func main() {
	var age int

	fmt.Print("Escribe tu edad: ")
	fmt.Scan(&age)

	if age < 18 {
		fmt.Println("No eres mayor de edad :(")
	} else if age == 18 {
		fmt.Println("Tienes justamente 18 años")
	} else {
		fmt.Println("Si eres mayor de edad :)")
	}
}
