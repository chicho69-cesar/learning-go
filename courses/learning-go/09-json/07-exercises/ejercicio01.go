/*
Comenzando con este código, marshal el slice []user a JSON. Hay una pequeña curva en
este ejercicio - recuerda preguntarte qué necesitas para EXPORTAR un valor de un paquete.
*/

package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name string `json:"name"`
	Age  int 	`json:"age"`
}

func main() {
	u1 := user{
		Name: "Cesar",
		Age:   21,
	}

	u2 := user{
		Name: "Alonso",
		Age:   23,
	}

	u3 := user{
		Name: "Liz",
		Age:   20,
	}

	users := []user {u1, u2, u3}
	fmt.Println(users)

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}
