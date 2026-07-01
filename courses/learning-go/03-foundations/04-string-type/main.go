/* 
“Un valor string es una (posiblemente vacía) secuencia de bytes. Los strings son
inmutables: una vez creados, es imposible cambiar su contenido.”
*/

package main

import "fmt"

func main() {
	s1 := "Hello World!!!"
	s2 := `Esta es una linea
	partida`

	fmt.Println(s1)
	fmt.Printf("El tipo de s1 es: %T\n", s1)

	fmt.Println(s2)
	fmt.Printf("El tipo de s1 es: %T\n", s2)

	fmt.Println("")

	bs := []byte(s1) // Convertimos el string a array de bytes o int8
	fmt.Println(bs)
	fmt.Printf("El tipo de bs es: %T\n", bs)

	fmt.Println("")

	for i := 0; i < len(s1); i++ {
		fmt.Printf("%#U ", s1[i]) // Imprimimos con #U el valor hexadecimal de cada caracter completo
	}

	fmt.Println("")

	for i, v := range s1 {
		fmt.Printf("En el indice %d el valor es %#x\n", i, v)
	}

	// El verbo %q nos permite imprimir otro verbo
	fmt.Printf("Con el verbo %q indico que se imprima el string %s", "%s", s1)
}
