package main

import "fmt"
import "mypackages/numbers"

func main() {
	var x uint = 41
	fmt.Printf("%d is even %v \n", x, numbers.Even(x))
}