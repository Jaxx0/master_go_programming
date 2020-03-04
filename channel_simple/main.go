package main

import "fmt"

func f1(n int, ch chan int)  {
	ch <- n
}

func main() {

	// Declaring and initilizing a RECEIVE-ONLY channel
	c3 := make(<-chan string)

	// Declaring and initialising a SEND-ONLY channel
	c4 := make(chan<- string)
	_,_ = c3,c4

	c := make(chan int)

	// receives
	c1 := make(<-chan string)

	//sends
	c2 := make(chan <- string)

	fmt.Printf("%T, %T, %T \n", c, c1, c2)

	go f1(20, c)
	n := <- c
	fmt.Println("Value received : ", n)
	fmt.Println("Exiting main ...")
}
