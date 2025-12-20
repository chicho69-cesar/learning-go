package main

import "fmt"

type person struct {
	name string
}

func main() {
	cesar := person {
		name: "Cesar",
	}

	fmt.Println(cesar) // name: "Cesar"
	change1(cesar)
	fmt.Println(cesar) // name: "Cesar"
	change2(&cesar)
	fmt.Println(cesar) // name: "Liz"
}

func change1(p person) {
	p.name = "Liz"
}

func change2(p *person) {
	p.name = "Liz"
}
