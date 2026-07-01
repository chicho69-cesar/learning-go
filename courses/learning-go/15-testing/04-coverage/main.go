/* 
Coverage (o cobertura) en programación es que tanto de nuestro código es cubierto por tests
(pruebas). Podemos usar el flag “-cover” para correr un análisis de cobertura en nuestro código.
Podemos usar el flag y el nombre del archivo requerido “-coverprofile <algún nombre de
archivo>” para escribir nuestro análisis de cobertura a un archivo.
código:
	# go test -cover
		go test -coverprofile c.out
			Mostrándolo en el navegador:
				go tool cover -html=c.out
			Aprender más
				go tool cover -h

Diferentes comandos útiles cuando hacemos benchmarks, ejemplos y tests.
	# godoc -http=:8080
	# go test
	# go test -bench .
		No olvides el “.” en la línea arriba
	# go test -cover
	# go test -coverprofile c.out
	# go tool cover -html=c.out
*/

package main

import "fmt"

func main() {
	fmt.Println(Greet("James"))
}

func Greet(s string) string {
	return fmt.Sprint("Hello my dear, ", s)
}
