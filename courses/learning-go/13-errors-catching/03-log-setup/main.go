package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
    log.SetOutput(file) // lanzamos nuestro archivo para historial de logs

	f, err := os.Open("sin-archivo.txt")
	if err != nil {
		log.Println("Ocurrio un error: ", err)
	}
	defer f.Close()

	fmt.Println("Revisa el archivo log.txt en el directorio")
}
