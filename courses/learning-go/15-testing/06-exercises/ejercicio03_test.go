package main

import "testing"

func TestCenteredAvg(t *testing.T) {
	type test struct {
		data   []int
		answer float64
	}

	tests := []test{
		{[]int{10, 20, 40, 60, 80}, 40},
		{[]int{2, 4, 6, 8, 10, 12}, 7},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5},
	}

	for _, test := range tests {
		got := CenteredAvg(test.data)

		if got != test.answer {
			t.Error("Got", got, "Want", test.answer)
		}
	}
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1000})
	}
}
