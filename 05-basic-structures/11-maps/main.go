/* 
Un mapa es un tipo de dato de almacenamiento llave-valor. Esto quiere decir que almacenas
algún valor y accedes al mismo con una llave. Por ejemplo, podría almacenar el valor
“numeroTel” y acceder a él con la llave “nombreAmigo”. De esta manera puedo ingresar el
nombre de mi amigo y el mapa me retornará su número telefónico. La sintaxis para el mapa es
map[llave]valor. La llave puede ser de cualquier tipo que permita comparación (por ejemplo,
podría usar un string o un int, ya que puedo comparar si dos strings son iguales o si dos
enteros son iguales. Es importante denotar que los mapas son desordenados. Si imprimes
todas las llaves y valores de un mapa se hará de forma aleatoria. El idioma comma ok es
también cubierto en este video. Un mapa es la estructura de datos perfecta cuando
necesitas buscar datos de manera rápida.
*/

package main

import "fmt"

func main() {
	m := map[string]int {
		"Batman": 32,
		"Robin":  27,
	}

	fmt.Println(m)

	fmt.Println(m["Batman"])
	fmt.Println(m["Cesar"]) /* Si el valor no esta en el mapa este regresara
	el valor zero por defecto del tipo de dato del valor del mapa */

	v, ok := m["Cesar"] // primer parámetro valor, segundo si esta o no en el mapa
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Cesar"]; ok {
		fmt.Println("Impresión desde el if", v)
	} else {
		fmt.Println("Impresión desde el else", v)
	}

	contacts := map[string]string {
		"Cesar": "3461099207",
		"Liz": "3461005286",
	}

	fmt.Println(contacts)

	contacts["Daniel"] = "4951890456"

	if value, ok := contacts["Daniel"]; ok {
		fmt.Printf("%s si esta en el mapa", value)
	}
}