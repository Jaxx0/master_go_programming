package main

import "fmt"

func f1()  {
	
}

func main() {
	var c1 chan float64

	// Declaring and initilizing a RECEIVE-ONLY channel
	c2 := make(<- chan rune)

	// Declaring and initilizing a SEND-ONLY channel
	c3 := make(chan <- rune)

	// a bidirectional buffered channel  called c4 with a capacity of 10 ints.
	c4 := make(chan int, 10)

	fmt.Printf("c1 :%T\n", c1)
	fmt.Printf("c2 :%#T\n", c2)
	fmt.Printf("c3 :%#T\n", c3)
	fmt.Printf("c4 :%#T\n", c4)
}