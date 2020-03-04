package main

import (
	"fmt"
	"os"
	"strconv"
)

func main()  {
	fmt.Println("od.Args", os.Args)
	fmt.Println("Path:", os.Args[0])
	fmt.Println("1st argument:", os.Args[1])
	//fmt.Println("2nd argument:", os.Args[2])
	//fmt.Println("3rd argument:", os.Args[3])
	//fmt.Println("No: of argument:", len(os.Args))
	var result, err  = strconv.ParseFloat(os.Args[1], 64)
	_ = err
	fmt.Printf("%T\n", os.Args[1])
	fmt.Printf("Type %T, value %f\n", result, result)
}
