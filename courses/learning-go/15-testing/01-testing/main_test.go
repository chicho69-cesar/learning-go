package main

import "testing"

func TestMySum(t *testing.T) {
	value := mySum(2, 8)
	if value != 10 {
		t.Error("Expected", 10, "Got", value)
	}
}
