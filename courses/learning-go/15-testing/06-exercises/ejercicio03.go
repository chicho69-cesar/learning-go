package main

import "sort"

func CenteredAvg(nums []int) float64 {
	sort.Ints(nums)
	nums = nums[1 : len(nums)-1]

	n := 0
	for _, v := range nums {
		n += v
	}

	f := float64(n) / float64(len(nums))
	return f
}
