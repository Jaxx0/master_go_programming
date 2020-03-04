package main

import "fmt"

var empty interface{}


func main() {
	fmt.Printf("%T\n", empty)
	empty = 5
	fmt.Printf("%T\n", empty)
	empty = 9.
	fmt.Printf("%T\n", empty)
	empty = []int{5,6,7,8,9}
	fmt.Printf("%T\n", empty)
	empty = append(empty.([]int), 3000)
	fmt.Printf("%T\n", empty)
	fmt.Printf("slice (empty variable) %v \n", empty)
}