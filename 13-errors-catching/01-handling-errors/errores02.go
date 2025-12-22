package main

import "fmt"

func main() {
	var response1, response2, response3 string

	fmt.Print("Nombre: ")
	_, err := fmt.Scan(&response1) // la función scan nos devuelve el numero de bytes y un error
	if err != nil {
        panic(err)
    }

	fmt.Print("Comida favorita: ")
	_, err = fmt.Scan(&response2)
    if err != nil {
		panic(err)
	}

	fmt.Print("Deporte favorito: ")
	_, err = fmt.Scan(&response3)
    if err != nil {
		panic(err)
	}

	fmt.Println(response1, response2, response3)
}
