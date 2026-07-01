package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello World!!!")
	
	/* aqui el metodo Fprintln recibe dos argumentos un elemento que
	implemente la interfaz writter y un string */
	fmt.Fprintln(os.Stdout, "Hello World!!!")

	/* este metodo es basicamente el mismo que el anterior */
	io.WriteString(os.Stdout, "Hello World!!!")
}
