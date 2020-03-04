package main

import "fmt"


func main() {
	c := make(chan string)

	go func (n string){
		c <- n
	}("mango")

	received :=  <- c
	fmt.Println("value received from channel is:", received)
}