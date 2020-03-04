package main

import (
	"fmt"
	"time"
)

func main()  {
	g()
}

func g()  {
	age := -9
	if age < 0 || age > 100 {
		fmt.Println("Invalid Age")
	} else if age <= 18 {
		fmt.Println("You are minor!")
	} else if age == 18 {
		fmt.Println("Congratulations! You've just become major!")
	} else {
		fmt.Println("You are major!")
	}
}

func f() {
	age := -9
	switch true {
	case age < 0 || age > 100:
		fmt.Println("Invalid Age")
	case age <= 18:
		fmt.Println("Invalid Age")
	case age <= 18 :
		fmt.Println("You are minor!")
	case  age == 18:
		fmt.Println("Congratulations! You've just become major!")
	default:
		fmt.Println("You are major!")
	}
}

func e()  {
	birthYear, currentYear := 1900, time.Now().Year()
	for year := birthYear; year <=currentYear;{
		fmt.Printf("Year %d  \n", year)
		year++
	}
}

func d()  {
	//print out all the numbers between 1 and 500 that divisible both by 7 and 5
	for i := 0; i <= 500; i++{
		if i%7 == 0 && i%5 == 0{
			fmt.Printf("%d is divisible by both 5 and 7 \n", i)
		}
	}
}

func c()  {
	count := 0
	for i :=0; i<=50; i++{
		if i%7 != 0{
			continue
		}

		fmt.Printf("%d is divisible by 7 \n", i)
		count++

		if count == 3{
			break
		}
	}
	fmt.Println("First statement after break")
}

func b()  {
	for i := 0; i <= 50; i++{
		if  i%7 != 0{
			continue
		}
		fmt.Printf("%d is divisible by 7 \n", i)
	}
}

func a()  {
	// print out all the numbers between 1 and 50 that divisible by 7
	for i := 0; i <50; i++{
		if i%7 == 0{
			fmt.Printf("%d is divisible by 7 \n", i)
		}

	}
}
