package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name 	 string `json:"name"`
	Lastname string `json:"lastname"`
	Age 	 int 	`json:"age"`
}

func main() {
	// creamos un string con formato json
	s := `[
		{"name":"Cesar","lastname":"Villalobos Olmos","age":21},
		{"name":"Liz","lastname":"Sandoval Vallejo","age":20}
	]`
	bs := []byte(s) // creamos un array de bytes con el string

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var persons []person
	err := json.Unmarshal(bs, &persons) // decodificamos y guardamos en persons
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Toda la data", persons)
	for i, v := range persons {
		fmt.Printf("\nLa persona numero: %d\n", i + 1)
		fmt.Println(v.Name, v.Lastname, v.Age)
	}
}
