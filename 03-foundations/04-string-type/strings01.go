package main

import "fmt"

func main() {
	s := "Hello, 世界"
	fmt.Printf("%s\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n\n", s)

	fmt.Printf("---%x\n\n", "世")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}
