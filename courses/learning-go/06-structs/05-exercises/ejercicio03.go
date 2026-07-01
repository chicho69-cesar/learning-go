/*
Crea un nuevo tipo: vehículo.
	El tipo subyacente es un struct.
	Los campos:
		puertas
		color
Crea dos nuevos tipos: camión & sedán.
	El tipo subyacente de cada uno de esos tipo es un struct.
	Embebe el tipo “vehículo” en ambos camión y sedán.
	Dale al camión el campo “cuatroRuedas” el cual será un booleano.
	Dale al sedán el campo “lujoso” el cual será un booleano.
Con los structs vehículo, camión y sedán:
	Usando un composite literal, crea un valor de tipo camión y asígnale valor a los campos.
	Usando un composite literal, crea un valor de tipo sedan y asígnale valor a los campos.
Imprime cada uno de los valores.
Imprime un solo valor de cada uno de eso valores.
*/

package main

import "fmt"

type vehiculo struct {
	puertas int
	color string
}

type camion struct {
	vehiculo
	cuatroRuedas bool
}

type sedan struct {
	vehiculo
	lujoso bool
}

func main() {
	c := camion {
		vehiculo: vehiculo {
			puertas: 4,
			color: "rojo",
		},
		cuatroRuedas: true,
	}

	s := sedan {
		vehiculo: vehiculo {
			puertas: 6,
			color: "negro",
		},
		lujoso: true,
	}

	fmt.Println(c)
	fmt.Println(s)

	fmt.Printf("Color del camion: %s\n", c.color)
	fmt.Printf("Color del sedan: %s\n", s.color)
}
