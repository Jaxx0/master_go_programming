package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f8()
}

func f8()  {
	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	newYears := append(years[:3], years[8:]...)
	fmt.Println(newYears)
}

func f7()  {
	friends := []string{"Marry", "John", "Paul", "Diana"}
	dst := make([]string, len(friends))
	nn := copy(dst, friends)
	fmt.Println(friends, dst, nn)
	dst[0] = "Perre"
	fmt.Println(friends, dst, nn)

}

func f6()  {
	nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}
	// Using a slice expression and a for loop iterate over the slice ignoring the first and the last two elements.
	sum := 0
	for _, num := range nums[2:6]{
		fmt.Println(num)
		sum += num
	}
	fmt.Printf("Sum: %d \n", sum)

}

func f5() {

	if len(os.Args) < 2 { //if not run with arguments
		log.Fatal("Please run the program with arguments (2-10 numbers)!")

	}

	//taking the arguments in a new slice
	args := os.Args[1:]

	// declaring and initializing sum and product of type float64
	sum, product := 0., 1.

	if len(args) < 2 || len(args) > 10 {
		fmt.Println("Please enter between 2 and 10 numbers!")
	} else {

		for _, v := range args {
			num, err := strconv.ParseFloat(v, 64)
			if err != nil {
				continue //if it can't convert string to float64, continue with the next number
			}
			sum += num
			product *= num
		}
	}

	fmt.Printf("Sum: %v, Product: %v\n", sum, product)

}

func f4() {
	//Create a Go program that reads some numbers from the command line and then calculates the sum and the product of all the numbers given at command line.
	//The user should give between 2 and 10 numbers.
	sum, product := 0., 1.
	for i := 1; i< len(os.Args); i++{
		num, _ := strconv.ParseFloat(os.Args[i], 64)
		fmt.Println(num)
		sum += num
		product *= num
	}
	fmt.Printf("Sum: %v, Product: %v \n", sum, product)

}

func f3() {
	nums := []float64{3.142, 2.99}
	nums = append(nums, 10.2)
	nums = append(nums, 4.1, 5.5, 6.6)
	fmt.Println(nums)

	n := []float64{2.2, 4.4}
	nums = append(nums, n...)
	fmt.Println(nums)
}

func f1() {
	nn := []string{"one", "two", "three"}
	for i, v := range nn {
		fmt.Printf("Index: %d, Value: %s \n", i, v)
	}
}

func f2() {
	mySlice := []float64{1.2, 5.6}

	mySlice[0] = 6.0

	a := 10
	mySlice[0] = float64(a)

	mySlice[1] = 10.10

	mySlice = append(mySlice, float64(a))

	fmt.Println(mySlice)

}
