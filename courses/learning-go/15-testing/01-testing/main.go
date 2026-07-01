package main

import "fmt"

func main() {
	fmt.Println("La suma de 2 + 4 es: ", mySum(2, 4))
	fmt.Println("La suma de 1 + 5 es: ", mySum(1, 5))
	fmt.Println("La suma de 6 + 7 es: ", mySum(6, 7))
}

func mySum(nums ...int) int {
	var sum int
    for _, num := range nums {
        sum += num
    }
    return sum
}
