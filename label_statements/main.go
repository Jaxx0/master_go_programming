package main

import "fmt"

func main()  {
	outer := 19
	_ = outer


	people := [5]string{"hellen", "mark", "brenda", "antonio", "michael"}
	friends := [2]string{"mark", "marry"}

outer:
	for index, name := range people{
		for _, friend := range friends{
			if name == friend{
				fmt.Printf("Found a friend %q at index %d \n", friend, index)
				break outer
			}
		}
	}
	fmt.Printf("Next instruction after the break")

}
