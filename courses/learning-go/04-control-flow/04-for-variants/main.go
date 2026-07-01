package main

import "fmt"

func main() {
	x := 1

	// PRIMEARA FORMA DE USAR FOR
	// for init; condición; post { }
	fmt.Println("\nAlternativa numero 1")
	for x = 1; x <= 10; x++ {
		fmt.Printf("\t\t1.- Valor de x: %d\n", x)
	}

	x = 1

	// SEGUNDA FORMA DE USAR FOR
	// for condición { }
	fmt.Println("\nAlternativa numero 2")
	for x <= 10 { // Alternativa a while (condicion)
		fmt.Printf("\t\t2.- Valor de x: %d\n", x)
		x++
	}

	x = 1

	// TERCERA FORMA DE USAR FOR
	// for
	fmt.Println("\nAlternativa numero 3")
	for { // Alternativa a while (true)
		if x > 10 {
			break
		}

		fmt.Printf("\t\t3.- Valor de x: %d\n", x)
		x++
	}
}
