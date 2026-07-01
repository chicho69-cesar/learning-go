/* 
Dentro de una declaración, el identificador pre-declarado iota representa una sucesión de
constantes enteras sin tipo. Es reiniciado a 0 cada vez que la palabra constante aparece en
el código. Puede ser usada para construir un conjunto de constantes relacionadas:
*/

package main

import (
	"fmt"
)

const (
	a = iota // Inicia valor constante en 0
	b // 1
	c // 2
)

const (
	d = iota // Reinicia el valor a 0, porque ya aparecia la keyword const
	e // 1
	f // 2
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
