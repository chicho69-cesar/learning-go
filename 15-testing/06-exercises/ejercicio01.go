/*
Comienza con este código. Obtenga el código listo para hacer BET en él (agregue benchmarks,
examples, tests). Ejecute lo siguiente en este orden:
● tests
● benchmarks
● coverage
● coverage mostrado en el navegador
● examples mostrados en documentación en el navegador web
*/

package main

func YearsOne(n int) int {
	return n * 7
}

func YearsTwo(n int) int {
	counter := 0
	for i := 0; i < n; i++ {
		counter += 7
	}
	return counter
}
