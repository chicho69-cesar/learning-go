/* 
Comienza con este código. En vez de usar el identificador blank (underscore), asegúrate de
que el código esté chequeando y manejando el error.
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person {
		First:   "James",
		Last:    "Bond",
		Sayings: []string { "Shaken, not stirred", "Any last wishes?", "Never say never" },
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(bs))
}
