/* 
Muy seguido, actualmente escuchamos acerca de bases de datos que han sido hackeadas y
los datos de los usuarios comprometidos. No hay excusas para esto. Nosotros, como
programadores, tenemos las herramientas para proteger la información de los usuarios. Bcrypt
es una de las herramientas que puedes usar para proteger los datos de los usuarios. Usando
bcrypt, obtendremos una mayor comprensión de cómo leer e implementar código de la
biblioteca estándar.
*/

package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`

	/* Aqui con el metodo GenerateFromPassword lo que hacemos es cifrar
	la contraseña que le pasamos como un array de bytes */
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)
	fmt.Println(bs)

	loginPword1 := `password123`

	/* Con el metodo CompareHashAndPassword lo que hacemos es comparar
	dos contraseñas diferentes */
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPword1))
	if err != nil {
		fmt.Println("YOU CAN'T LOGIN")
		return
	}

	fmt.Println("You're logged in")
}
