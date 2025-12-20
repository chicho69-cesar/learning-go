/* 
Comenzando con este código, unmarshal el JSON a un estructura de datos de Go. Pista:
https://mholt.github.io/json-to-go/
*/

package main

import (
	"fmt"
	"encoding/json"
)

type Usuario struct {
	Nombre 	 string
	Apellido string
	Edad 	 int
	Dichos 	 []string
}

func main() {
	s := `[{"Nombre":"Eduar","Apellido":"Tua","Edad":32,"Dichos":["Cachicamo diciéndole a morrocoy conchudo","La mona, aunque se vista de seda, mona se queda","Poco a poco se anda lejos"]},{"Nombre":"Carlos","Apellido":"Pérez","Edad":27,"Dichos":["Camarón que se duerme se lo lleva la corriente","A ponerse las alpargatas que lo que viene es joropo","No gastes pólvora en zamuro"]},{"Nombre":"M","Apellido":"Hmmmm","Edad":54,"Dichos":["Ni lava ni presta la batea","Hijo de gato, caza ratón","Más vale pájaro en mano que cien volando"]}]`
	fmt.Println(s)

	bs := []byte(s)
	fmt.Println(bs)

	var usuarios []Usuario
	err := json.Unmarshal(bs, &usuarios)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("TODA LA INFORMACIÓN")
	for _, usuario := range usuarios {
		fmt.Println(usuario.Nombre, usuario.Apellido, usuario.Edad)

		for _, dicho := range usuario.Dichos {
			fmt.Println("\t", dicho)
		}
	}
}
