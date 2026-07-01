package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// creamos un archivo de texto
	file, err := os.Create("names.txt")
	if err != nil {
        fmt.Println(err)
		return
    }
	defer file.Close() // la función defer se ejecuta al retornar en la función donde se define

	row := strings.NewReader("Cesar Villalobos Olmos")

	io.Copy(file, row) // agregamos el row al archivo de texto
}
