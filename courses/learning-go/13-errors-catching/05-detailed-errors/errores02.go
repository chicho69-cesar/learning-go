package main

import (
	"fmt"
	"math"
	"log"
)

type ubicationError struct {
	lat  string
	long string
	err  error
}

// .todos los tipos que implementan el metodo Error se convierten en Exceptions
func (u ubicationError) Error() string {
	return fmt.Sprintf("Un error de concepto matematico a ocurrido: %v %v %v", u.lat, u.long, u.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		name := fmt.Errorf("Error mat: raíz cuadrada de numero negativo: %v", f)
		return 0, ubicationError { "50.2289 N", "99.4656 W", name }
	}

	return math.Sqrt(f), nil
}
