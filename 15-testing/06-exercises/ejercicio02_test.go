package main

import "testing"

func TestCount(t *testing.T) {
	n := Count("one two three")
	if n != 3 {
		t.Error("Got", n, "Wanted", 3)
	}
}

func TestUseCount(t *testing.T) {
	n := UseCount("one two three three three")

	for k, v := range n {
		switch k {
		case "one":
			if v != 1 {
				t.Error("Got", v, "Wanted", 1)
			}
		case "two":
			if v != 1 {
				t.Error("Got", v, "Wanted", 1)
			}
		case "three":
			if v != 3 {
				t.Error("Got", v, "Wanted", 3)
			}
		}
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count("one two three")
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount("one two three three three")
	}
}
