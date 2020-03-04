package main

import "fmt"


func main() {
	ch := make(chan int)

	for i:=1; i<51; i++{
		go func(i int){
			ans := i*i
			ch <- ans
		}(i)
		fmt.Printf("Square of %v is %v \n", i, <- ch)
	}
}