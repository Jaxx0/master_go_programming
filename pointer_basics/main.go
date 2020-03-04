package main

import "fmt"

func main()  {
	name := "Andrei"
	fmt.Println(&name)
	fmt.Println(*&name)

	var x int = 5
	ptr := &x
	fmt.Printf("ptr is of type %T and value %v and address %p\n", ptr, ptr, &ptr)
	fmt.Printf("The address of x is %p \n", &x)

	var ptr1 *float64
	_ = ptr1

	p := new(int)
	x = 100
	p = &x
	fmt.Printf("P is of type %T and value %v \n", p, p)
	fmt.Printf("The address of x is %p \n", &x)

	*p = 90 //equivalent to x = 90
	fmt.Println(x, *p)   // de-referencing operator
	fmt.Println("*p == x", *p==x)

	*p = 10
	*p = *p/2
	fmt.Println(x)
}
