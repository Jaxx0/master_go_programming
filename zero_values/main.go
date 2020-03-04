package main

import "fmt"

func main()  {
	var a int
	var b float64

	a = 5
	b = 5.5

	a = int(b)
	fmt.Println(a,b)

	var value int
	var price float64
	var name  string
	var done bool
	var ppt = &value

	fmt.Println(value, price, name, done, ppt)

	var x float64
	fmt.Println(x == 0.0)


}
