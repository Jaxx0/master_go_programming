package main

import (
	"fmt"
	calcv1 "github.com/Jaxx0/go_math/calc"
	"github.com/Jaxx0/go_math/geometry"
	calcv2 "github.com/Jaxx0/go_math/v2/calc"
)

func main()  {
	fmt.Println(geometry.CubeVolume(0))
	fmt.Println(calcv1.Add(20,45))
	fmt.Println(calcv1.Sub(500,250))
	s := calcv2.Add(4,5,6)
	fmt.Println(s)
}
