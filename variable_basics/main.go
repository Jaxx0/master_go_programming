package main

import "fmt"

func main()  {
	var age = 30
	fmt.Println("Age:", age)

	var name = "Jackson"
	//fmt.Println("Name is:", name)

	_ = name

	s := "Learning GO lang"
	fmt.Println("s:", s)

	name = "Alabama"
	song1 := "Sweet Home"
	_ = song1

	car, cost := "Audi", 40000
	fmt.Println(car, cost)

	var i,j int
	i,j = 5, 8

	fmt.Println(j,i)
	sum := 5 + 2.5
	fmt.Println(sum)

	//_ blank variable avoids error of declared but not used


}
