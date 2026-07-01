/* 
Un “callback” es cuando le pasamos una función a otra función como argumento. Para este
ejercicio,
	Pasa una func a otra función como argumento
*/

package main

import "fmt"

func main() {
	g := func (nums []int) int {
		if len(nums) == 0 {
			return 0
		}

		if len(nums) == 1 {
			return nums[0]
		}

		return nums[0] + nums[len(nums) - 1]
	}

	nums := []int { 1, 2, 3, 4, 5, 6, 7 }
	fmt.Println(foo(g, nums))
}

func foo(f func(nums []int) int, nums []int) int {
	result := f(nums)
	result++
	return result
}
