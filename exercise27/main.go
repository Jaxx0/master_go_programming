package main

import (
	"fmt"
	"github.com/Jaxx0/go_arithmetic/IsPrime"
	"github.com/Jaxx0/go_arithmetic/factorial"
)

func main() {
	var x int = 3
	fmt.Printf("%d is prime: %v \n",x, IsPrime.IsPrime(x))
	fmt.Printf("%d Factorial is: %v \n",x, factorial.Factorial(x))
}