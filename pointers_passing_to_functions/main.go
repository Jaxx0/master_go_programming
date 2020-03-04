package main

import "fmt"

func change(a *int) *float64 {
	*a = 100
	b := 5.5
	return &b
}

func changeVar(a int)  {
	a = 66
}

func main()  {
	x := 8
	p := &x
	fmt.Printf("Value of x before changing it %v \n", x)
	change(p)
	fmt.Printf("Value of x after changing it %v \n", x)
	fmt.Printf("Value of x before calling changevar() is %v \n", x)
	changeVar(x)
	fmt.Printf("Value of x after calling changevar() is %v \n", x)

}
