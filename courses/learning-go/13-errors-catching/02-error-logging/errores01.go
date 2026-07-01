package main

import (
	"fmt"
	// "log"
	"os"
)

func main() {
	defer foo()

	_, err := os.Open("sin-archivo.txt")
	if err != nil {
		// fmt.Println(err) // imprimir con Println
		// log.Println(err) // imprimir con log.Println - muestra fecha y hora
		// log.Panicln(err) // nos manda salir con error
		// log.Fatalln(err) // muestra el codigo de status de la salida
		panic(err) // rompe la ejecución con error
	}
}

func foo() {
	fmt.Println("Cuando os.Exit() es llamada, las funciones diferidas no corren")
}
