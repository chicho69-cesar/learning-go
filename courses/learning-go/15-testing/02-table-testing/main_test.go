package main

import "testing"

func TestMySum(t *testing.T) {
	type test struct {
		data     []int
		response int
	}

	tests := []test{
		test{[]int{2, 4, 6}, 12},
		test{[]int{1, 5, 2}, 8},
		test{[]int{0, -1, 1}, 0},
		test{[]int{-10, 4, 20}, 14},
		test{[]int{1, 2, 66}, 69},
	}

	for _, test := range tests {
		value := mySum(test.data...)
		if value != test.response {
			t.Error("Expected", test.response, "Got", value)
		}
	}
}
