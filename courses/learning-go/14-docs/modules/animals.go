package main

import "fmt"

type Animal struct {
	Name string
	Race string
}

func (a Animal) Info() string {
	return fmt.Sprintf("%s: %s", a.Name, a.Race)
}

type Dog struct {
	Animal
	Weight float64
}

func (d Dog) Info() string {
	return fmt.Sprintf("%s: %s, %g", d.Name, d.Race, d.Weight)
}

type Cat struct {
    Animal
	Lives int
}

func (c Cat) Info() string {
	return fmt.Sprintf("%s: %s, %d", c.Name, c.Race, c.Lives)
}

type IAnimal interface {
	Info() string
}

func ShowInfo(a IAnimal) string {
	return a.Info()
}
