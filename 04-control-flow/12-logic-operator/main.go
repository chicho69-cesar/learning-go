/* 
Los operadores logicos se usan para construir expresiones logicas que nos devuelven valores booleanos
Estos son:
    And && (Conjuncion) => and => Multiplicacion logica
    Or || (Disyuncion) => or => Suma logica
    Negacicon ! => not

And -> tabla de resultados
    true and true -> true
    true and false -> false
    false and true -> false
    false and false -> false

Or -> tabla de resultados
	true or true -> true
	true or false -> true
	false or true -> true
	false or false -> false

Not -> tabla de resultados
    not(true) -> false
    not(false) -> true

Prioridad de los operadores logicos
    1.- Not
    2.- And
    3.- Or

Prioridad global de los operadores en general
    1.- ()
    2.- **
    3.- *, /, mod(%), !
    4.- +, -, &&
    5.- <, >, ==, <=, >=, !=, ||
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(true && true) // true
	fmt.Println(true && false) // false
	fmt.Println(true || true) // true
	fmt.Println(true || false) // true
	fmt.Println(!true) // false
}
