package main

import "fmt"

type person struct {
	name     string
	lastname string
}

type singer struct {
	person
	genders []string
}

func (p person) speak() {
	fmt.Println("La persona ", p.name, p.lastname, " esta hablando")
}

func (s singer) speak() {
	fmt.Println("La persona ", s.name, s.lastname, " esta cantando")
}

type human interface {
	speak()
}

func bar(h human) {
	/* Cuando tenemos una interfaz podemos acceder al tipo que la esta
	implementando usando h.(type) */
	switch h.(type) {
		case person:
			/* Cuanod estamos seguros del tipo que implementa la interfaz
			podemos hacer uncasting a ese tipo de la interfaz usando:
			h.(type).field */
			fmt.Println("I'm a person passed into bar ", h.(person).name)
		case singer:
			fmt.Println("I'm a singer passed into bar ", h.(singer).name)
	}

	fmt.Println("I'm also a human passed into bar ", h)
}

func main() {
	s1 := singer {
		person: person {
			name: "Cesar",
			lastname: "Villalobos Olmos",
		},
		genders: []string {
			"Rock",
			"Hip Hop",
			"Pop",
		},
	}

	s2 := singer {
		person: person {
			name: "Liz",
			lastname: "Sandoval Vallejo",
		},
		genders: []string {
			"Hip Hop",
			"Pop",
			"Reggaeton",
		},
	}

	p := person {
		name: "Danna",
		lastname: "Delgado Hernandez",
	}

	fmt.Println(s1)
	fmt.Println(s2)

	s1.speak()
	s2.speak()
	p.speak()

	bar(s1)
	bar(s2)
	bar(p)
}
