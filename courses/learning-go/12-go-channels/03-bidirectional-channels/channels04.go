package main

import "fmt"

func main() {
	c := make(chan string)
	cr := make(<-chan string)
	cs := make(chan<- string)

	fmt.Printf("Channel Bidirectional: \t%T\n", c)
	fmt.Printf("Channel Recieve Only: \t%T\n", cr)
	fmt.Printf("Channel Send Only: \t%T\n", cs)
}
