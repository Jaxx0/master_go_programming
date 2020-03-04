package main

import (
	"fmt"
	"math"
)

func f1(a int, b int)  {
	fmt.Println("Sum:", a+b)
}

func f2(a,b,c int, d,e float64, s string)  {
	fmt.Println(a,b,c,d,e,s)
}

func f3(a float64) float64 {
	p:= math.Pow(a, a)
	fmt.Println(p)
	return p
}

func f4(a,b int) (int, int)  {
	return a+b, a*b
}

func f5(a,b int) (s int) {
	fmt.Println(s)
	s = a +b
	return
}

func main()  {
	f1(2,5)
	f2(4,5,6,7.7,8.8,"golang")
	f3(5.1)
	a, b := f4(8,9)
	fmt.Println(a,b)
	mySum := f5(4,8)
	fmt.Println(mySum)
}