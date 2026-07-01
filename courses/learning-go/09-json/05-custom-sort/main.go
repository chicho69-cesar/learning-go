package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

type ByAge []Person
type ByName []Person

// creamos los metodos necesarios para que el paquete sort puede ordenar en base a las condiciones aqui definidas
func (a ByAge) Len() int { return len(a) }
func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] } 
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// estos metodos hacen que el tipo ByName pertenezca a Interface
func (n ByName) Len() int { return len(n) }
func (n ByName) Swap(i, j int) { n[i], n[j] = n[j], n[i] }
func (n ByName) Less(i, j int) bool { return n[i].Name < n[j].Name }

func main() {
	p1 := Person{
		Name: "Cesar",
		Age:  21,
	}

	p2 := Person{
		Name: "Alonso",
		Age:  23,
	}

	p3 := Person{
		Name: "Liz",
		Age:  20,
	}

	p4 := Person{
		Name: "Yuly",
		Age:  18,
	}

	persons1 := []Person{p1, p2, p3, p4}
	persons2 := []Person{p1, p2, p3, p4}

	fmt.Println(persons1)
	sort.Sort(ByAge(persons1)) // el metodo sort recibe un tipo Interface
	fmt.Println(persons1)

	fmt.Println(persons2)
	sort.Sort(ByName(persons2)) // ordenamos por Name
	fmt.Println(persons2)
}
