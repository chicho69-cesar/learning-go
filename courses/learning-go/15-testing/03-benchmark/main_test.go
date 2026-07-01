package main

import "testing"

func TestSaludar(t *testing.T) {
	s := Saludar("Cesar")
	if s != "Bienvenido querido Cesar" {
		t.Error("Expected", "Bienvenido querido Cesar", "Got", s)
	}
}

/* Esta es la función Benchmark que nos va a ayudar a medir el rendimiento
de nuestro codigo */
func BenchmarkSaludar(b *testing.B) {
	// El valor N es el numero de veces que el paquete determina
	for i := 0; i < b.N; i++ {
		Saludar("Cesar")
	}
}
