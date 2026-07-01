package main

import "fmt"

var x int

func main() {
	// if - else
	x = 42
	if x == 40 {
		fmt.Println("El valor es 40")
	} else {
		fmt.Println("El valor no es 40")
	}


	// if - else if - else
	x = 434
	if x == 40 {
		fmt.Println("El valor es 40")
	} else if x == 41 {
		fmt.Println("El valor es 41")
	} else {
		fmt.Println("El valor no es 40 ni 41")
	}


	// if - else if ... - else 
	x = 434
	if x == 40 {
		fmt.Println("El valor es 40")
	} else if x == 41 {
		fmt.Println("El valor es 41")
	} else if x == 42 {
		fmt.Println("El valor es 42")
	} else if x == 43 {
		fmt.Println("El valor es 43")
	} else {
		fmt.Println("El valor no es 40, 41, 42, ni 43")
	}
}
