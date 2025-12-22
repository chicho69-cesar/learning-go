/* 
Comienza con este código. Crea un mensaje de error personalizado usando “fmt.Errorf”.
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type persona struct {
	Nombre   string
	Apellido string
	Frases   []string
}

func main() {
	p1 := persona {
		Nombre:   "James",
		Apellido: "Bond",
		Frases:   []string{ "Shaken, not stirred", "¿Algún último deseo?", "Nunca digas nunca." },
	}

	bs, err := aJSON(p1)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(bs))
}

// aJSON también necesita retorna un error
func aJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
        return []byte {}, fmt.Errorf("Ocurrio un error al decodificar el JSON: %v", err)
    }

	return bs, nil
}
