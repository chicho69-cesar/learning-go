package main

import "fmt"

func main() {
	x := 0
	foo(x)
	fmt.Println(x) //0
}

func foo(y int) {
	fmt.Println(y) // 0
	y = 21
	fmt.Println(y) // 21
}
