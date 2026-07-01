package main

import (
	"fmt"
)

var a int
var b string = "James Bond"

func main() {
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\t%T\n", a, b)

	s := fmt.Sprint(a, " algo más ", b)
	fmt.Println(s)
	
	s2 := fmt.Sprintf("%v\t%T\t%T\n", "valores a pasar", a, b)
	fmt.Println(s2)
}
