package main

import "fmt"

func main()  {
	// %q - quoted string
	var figure string = "circle"
	fmt.Printf("This is a %s \n", figure)
	fmt.Printf("This is a %q \n", figure)

	//%v replaced by any type
	fmt.Printf("This is a %v \n", figure)

	// %t bool
	closed := true
	fmt.Printf("File closed? %t \n", closed)

	//%b - base 2
	fmt.Printf("%b \n", 55)

	//ip addresses 8 bits
	fmt.Printf("%08b \n", 60)

	//%x - hex
	fmt.Printf("%x \n", 100)

	r := 3.4
	t := 6.9

	fmt.Printf("r * t = %f \n", r*t)
	fmt.Printf("r * t = %.3f \n", r*t)
}
