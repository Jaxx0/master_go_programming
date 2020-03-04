package main

import "fmt"
// declare an array

var numbers [10]int

// initialise global variables

func init()  {
	fmt.Println("This is init 1")
}

func init()  {
	fmt.Println("This is init 2")
}

func init()  {
	fmt.Println("This is init 3")
}

func main()  {
	fmt.Println("This is main")
	//init()
	fmt.Printf("%#v \n", numbers)
}
