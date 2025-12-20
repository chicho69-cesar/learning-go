/* Comenzando con este código, ordena el []usuario por
	# Edad
	# Apellido
También ordena el []string “Dichos” para cada usuario
	# Imprime todo de una manera agradable
*/

package main

import (
	"fmt"
	"sort"
)

type Usuario struct {
	Nombre   string
	Apellido string
	Edad     int
	Dichos   []string
}

type PorApellido []Usuario
type PorEdad []Usuario

func (pn PorApellido) Len() int { return len(pn) }
func (pn PorApellido) Swap(i, j int) { pn[i], pn[j] = pn[j], pn[i] }
func (pn PorApellido) Less(i, j int) bool { return pn[i].Apellido < pn[j].Apellido }

func (pe PorEdad) Len() int { return len(pe) }
func (pe PorEdad) Swap(i, j int) { pe[i], pe[j] = pe[j], pe[i] }
func (pe PorEdad) Less(i, j int) bool { return pe[i].Edad < pe[j].Edad }

func main() {
	u1 := Usuario{
		Nombre:   "Eduar",
		Apellido: "Tua",
		Edad:     32,
		Dichos: []string{
			"Cachicamo diciéndole a morrocoy conchudo",
			"La mona, aunque se vista de seda, mona se queda",
			"Poco a poco se anda lejos",
		},
	}

	u2 := Usuario{
		Nombre:   "Carlos",
		Apellido: "Pérez",
		Edad:     27,
		Dichos: []string{
			"Camarón que se duerme se lo lleva la corriente",
			"A ponerse las alpargatas que lo que viene es joropo",
			"No gastes pólvora en zamuro",
		},
	}

	u3 := Usuario{
		Nombre:   "Che",
		Apellido: "López",
		Edad:     54,
		Dichos: []string{
			"Ni lava ni presta la batea",
			"Hijo de gato, caza ratón",
			"Más vale pájaro en mano que cien volando",
		},
	}

	usuarios := []Usuario{u1, u2, u3}
	fmt.Println(usuarios)

	sort.Sort(PorEdad(usuarios))

	fmt.Println("\nPOR EDAD")
	for _, usuario := range usuarios {
		sort.Strings(usuario.Dichos)
		fmt.Printf("Nombre: %s\tApellido: %s\tEdad: %d años\n", usuario.Nombre, usuario.Apellido, usuario.Edad)

		for _, dicho := range usuario.Dichos {
			fmt.Printf("\t%s\n", dicho)
		}
	}

	sort.Sort(PorApellido(usuarios))

	fmt.Println("\nPOR APELLIDOS")
	for _, usuario := range usuarios {
		sort.Strings(usuario.Dichos)
		fmt.Printf("Nombre: %s\tApellido: %s\tEdad: %d años\n", usuario.Nombre, usuario.Apellido, usuario.Edad)

		for _, dicho := range usuario.Dichos {
			fmt.Printf("\t%s\n", dicho)
		}
	}
}
