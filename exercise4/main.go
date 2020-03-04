package main

import "fmt"

func main()  {
	f3()
}


func f3() {
	myArray := [3]float64{1.2, 5.6}

	myArray[2] = 6

	a := 10
	myArray[0] = float64(a)

	myArray[2] = 10.10

	fmt.Println(myArray)

}


//Write a Go program that counts the number of positive even numbers in the array
func f2()  {
	nums := []int{30, -1, -6, 90, -6}
	count := 0
	for _,v := range nums{
		if v > 0 && v%2 ==0{
			count ++
			fmt.Println(v)
		}
	}
	fmt.Printf("We have %d positive even numbers", count)
}

func f1()  {
	var cities [2]string
	fmt.Printf("%#v\n", cities)
	fmt.Println(cities)

	grades := [3]float64{0:82, 1:69}
	fmt.Println(grades)

	taskDone := [...]bool{}
	_ =taskDone

	for i := 0; i < len(cities); i++{
		fmt.Printf("%q\n", cities[i])
	}

	for i,v := range cities{
		fmt.Printf("Index %d, value:%q\n", i, v)
	}

	for _,v := range grades{
		fmt.Println(v)
	}

}