package main

import (
	"fmt"
)

func main()  {
	i := 0
loop:
	if i < 5{
		fmt.Println(i)
		i++
		goto loop
	}

}


func a() {
	//goto todo
	x := 10
	fmt.Println(x)
	goto todo


todo:
	fmt.Println("something here")
}