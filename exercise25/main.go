package main

import "fmt"

func power(n int, ch chan int)  {
	square := n*n
	ch <- square
}

func main() {
	ch := make(chan int)

	for i:=1; i<51; i++{
		go power(i, ch)
		fmt.Printf("Square of %v is %v \n", i, <- ch)
	}
}