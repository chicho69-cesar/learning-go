package main

import "testing"

func TestYearsOne(t *testing.T) {
	result := YearsOne(10)
	if result != 70 {
		t.Error("Got", result, "Want", 70)
	}
}

func TestYearsTwo(t *testing.T) {
	result := YearsTwo(10)
	if result != 70 {
		t.Error("Got", result, "Want", 70)
	}
}

func BenchmarkYearsOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsOne(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}
