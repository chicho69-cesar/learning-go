package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World!")
}

func Cat(strs []string) string {
	str := strs[0]
	for _, s := range strs[1:] {
		str += " "
		str += s
	}
	return str
}

func Join(strs []string) string {
	return strings.Join(strs, " ")
}
