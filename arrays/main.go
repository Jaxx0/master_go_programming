package main

import (
	"fmt"
	"strings"
)

func main()  {
	var numbers [3]int
	fmt.Printf("%v\n", numbers)
	fmt.Printf("%#v\n", numbers)

	var val = [4]float64{}
	fmt.Printf("%#v\n", val)

	var a1 = [3]int{1,4,-10}
	fmt.Printf("%#v\n", a1)

	a2 := [4]string{"Dan", "Diana", "Paul", "John"}
	fmt.Printf("%#v\n", a2)

	a4 := [4]string{"x", "y"}
	fmt.Printf("%#v\n", a4)

	// ellipsis operator
	a5 := [...]int{1,2,3,4,5,6,7,5,5,6,7,8,7,9,9,9}
	fmt.Printf("%#v\n", a5)
	fmt.Printf("The length of a5 is %d\n", len(a5))

	a6 := [...]int{1,
		2,
		3,
		4,
		5,
		6,
		7,}
	fmt.Printf("%#v\n", a6)

	numbers[0] = 7
	fmt.Println(numbers)

	numbers[2] = 100
	fmt.Println(numbers)

	for i,v := range numbers{
		fmt.Println("index:", i, "value:", v)
	}

	fmt.Println(strings.Repeat("#", 20))

	for i:= 0; i<len(numbers); i++{
		fmt.Println("index:", i, "value:", numbers[i])
	}

	balances := [2][3]int{
		{5,6,7},
		[3]int{8,9,10},
		}
	fmt.Printf("%v\n", balances)


	// array copies are independent and unique
	m := [3]int{1,2,3}
	n := m
	fmt.Println("n is equal m:", n==m)

	// modify m
	m[0] = -1
	fmt.Println("n is equal m:", n==m)

	fmt.Println(strings.Repeat("_", 50))

	grades := [3]int{
		1:10,
		0:5,
		2:7}
	fmt.Println(grades)

	accounts := [2]int{1:50}
	fmt.Println(accounts)

	names := [...]string{5:"Duncan"}
	fmt.Println(names, len(names))

	integers := [...]int{5:5}
	fmt.Println(integers, len(integers))

	cities := [...]string{5:"Paris", "London", 1: "NYC"}
	fmt.Println(cities, len(cities))

	for i,v := range cities{
		fmt.Printf("%d - %q \n", i, v)
	}

	fmt.Printf("%#v\n", cities)

	weekend := [7]bool{5:true, 6:true}
	fmt.Printf("%#v\n", weekend)

}
