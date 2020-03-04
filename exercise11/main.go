package main

import "fmt"

func f1()  {
	x := 10.10
	ptr := &x
	fmt.Printf("The address of x is: %p \n",&x)
	fmt.Printf("Type: %T, Value: %v \n", ptr, ptr)
	fmt.Printf("Address of pointer: %p, Value of x: %v \n", &ptr, *ptr)
}

func swapFloatsByPointer(x, y *float64) {
	*x, *y = *y, *x
}

func divide()  {
	x, y := 10, 2
	ptrx, ptry := &x, &y
	z := (*ptrx)/(*ptry)
	fmt.Println(z)
}

func main() {
	x, y := 5.5, 8.8
	fmt.Printf("BEFORE swapping x: %v, y: %v \n", x, y)
	swapFloatsByPointer(&x, &y)
	fmt.Printf("AFTER swapping x: %v, y: %v \n", x, y)
	divide()
}

