/*
Comenzando con este código, usa el struct raiz.Error como valor de tipo error. Si quieres usa
estos valores para tu
● lat "50.2289 N"
● long "99.4656 W"
*/

package main

import (
	"fmt"
	"log"
	"math"
)

type raizError struct {
	lat  string
	long string
	err  error
}

func (re raizError) Error() string {
	return fmt.Sprintf("Error matemático: %v %v %v", re.lat, re.long, re.err)
}

func main() {
	_, err := raiz(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func raiz(f float64) (float64, error) {
	if f < 0 {
		err := raizError {
			lat:  "50.2289 N",
            long: "99.4656 W",
            err:  fmt.Errorf("Error matemático"),
		}

		return 0, err
	}

	return math.Sqrt(f), nil
}
