package main

import (
	"encoding/json"
	"fmt"
)

/* Aqui es super importante poner los campos en mayusculas, ya que de esta
forma declaramos que los campos van a ser visibles para las funciones de 
otros paquetes, en este caso las del json y en el tercer valor especificamos
el nombre que va a tener el campo en el json */
type person struct {
	Name 	 string `json:"name"`
	Lastname string `json:"lastname"`
	Age 	 int 	`json:"age"`
}

func main() {
	p1 := person {
		Name: "Cesar",
		Lastname: "Villalobos Olmos",
		Age: 21,
	}

	p2 := person {
		Name: "Liz",
		Lastname: "Sandoval Vallejo",
		Age: 20,
	}

	persons := []person { p1, p2 }
	fmt.Println(persons)

	bs, err := json.Marshal(persons)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
