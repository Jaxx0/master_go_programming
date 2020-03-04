package main

import (
	"fmt"
)

func main()  {
	const days int  = 7

	// in a group constant dubsequent constants pick value from preceding one
	const (
		min1 = -500
		min2
		min3
	)

	fmt.Println(min1, min2, min3)
	//const power = math.Pow(2, 3)  // Power is a runtime function
	//fmt.Println(power)

	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)

	fmt.Println(c1, c2, c3)

	const (
		c11 = iota
		c22
		c33
	)
	fmt.Println(c11, c22, c33)

	const (
		c111 = 777
		c222
		c333
	)
	fmt.Println(c111, c222, c333)

	const (
		North = iota
		East
		South
		West
	)
	fmt.Println(North, East, South, West)

	const (
		a = iota * 3
		b
		c
		d
	)
	fmt.Println(a, b, c, d)

	const (
		x = -(iota + 2) //-2
		_     // skip the next i.e -3
		y				// -3
		z				// -4
		zz				// -5
	)
	fmt.Println(x, y, z, zz)

	const p = 7  // untyped constant - will get type from the first expression
	const q float64 = 3.1
	fmt.Println(p * q)

	//There are ONLY boolean constants,
	//rune constants, integer constants,
	//floating-point constants, complex constants,
	//and string constants.

}
