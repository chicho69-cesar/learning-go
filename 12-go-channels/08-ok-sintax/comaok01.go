package main

import "fmt"

func main() {
	ch := make(chan int)
	
	go func() {
        ch <- 21
		close(ch)
	}()

	v, ok := <-ch // si el canal no esta vacio toma el valor que tiene y true
	fmt.Println(v, ok)

	v, ok = <-ch // si el canal esta vacio toma los valores cero por default, 0 y false
	fmt.Println(v, ok)
}
