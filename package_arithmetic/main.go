package main

import "fmt"
import "mypackages/arithmetic"

func main()  {
	var x int = 13
	fmt.Printf("%d is prime: %v \n", x, arithmetic.IsPrime(x))
}
