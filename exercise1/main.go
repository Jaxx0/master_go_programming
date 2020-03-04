package main

import (
	"fmt"
	"strconv"
)

func main() {
	u()
}

type mile float64
type kilometer float64

const m2km = 1.609

func u() {
	var mileBerlinToParis mile = 655.3
	var kmBerlinToParis kilometer
	kmBerlinToParis = kilometer(mileBerlinToParis * m2km)
	fmt.Printf("Distance fom Paris to Berlin is %f kilometers \n", kmBerlinToParis)
}

type duration int

func t() {
	var hour duration = 3600
	minute := 60
	fmt.Println(hour != duration(minute))
}

var hour duration

func s() {
	fmt.Println(hour)
	fmt.Printf("Type %T\n", hour)
	hour = 3600
	fmt.Printf("Value %v\n", hour)
}

func r() {
	x, y := 0.1, 5
	var z float64

	// Write the correct logical operator (&&, || , !)
	// inside the expression so that result1 will be false and result2 will be true.

	result1 := (x < z) || (int(x) != int(z))
	fmt.Println(result1)

	result2 := y == (1*5) && (int(z) == 0)
	fmt.Println(result2)

}
func q() {
	const distanceFromSun = 149_600_000 * 1000
	const speedOfLight = 299_792_458
	fmt.Printf("It takes %d seconds for sunlight to hit the earth", distanceFromSun/speedOfLight)
}

func p() {
	x, y := 4, 5.1

	z := float64(x) * y
	fmt.Println(z)

	a := 5
	b := 6.2 * float64(a)
	fmt.Println(b)

}

func o() {
	var i = 3
	var f = 3.2
	var s1, s2 = "3.14", "5"

	//  i to string (int to string)
	var j = strconv.Itoa(i)
	fmt.Printf("Type:%T - Value: %q\n", j, j)

	//s2 to int (string to int)
	var i2, _ = strconv.Atoi(s2)
	fmt.Printf("Type:%T - Value: %d\n", i2, i2)

	// f to string (float64 to string)
	fs := fmt.Sprintf("%f", f)
	fmt.Printf("Type:%T - Value: %q\n", fs, fs)

	// 3. float64 to string
	ss1 := fmt.Sprintf("%f", f)
	fmt.Printf("ss1's type: %T, ss1's value: %q\n", ss1, ss1)

	// s1 to float64  (string to float64)
	iff, _ := strconv.ParseFloat(s1, 64)
	fmt.Printf("Type:%T - Value: %f\n", iff, iff)

}

func n() {
	var i int = 3
	var f float64 = 3.2

	// convert i to float64
	fmt.Printf("i in float64 is %f\n", float64(i))

	// convert f to int
	fmt.Printf("f in int is %d\n", int(f))

	fmt.Printf("Type: %T - Value: %f\n", float64(i), float64(i))
	fmt.Printf("Type: %T - Value: %d\n", int(f), int(f))

}

func m() {
	shape := "circle"
	radius := 3.2
	const pi float64 = 3.14159
	circumference := 2 * pi * radius

	fmt.Printf("Shape: %q\n", shape)
	fmt.Printf("Circle's circumference with radius %f is %f\n", radius, circumference)
}

func l() {
	const x float64 = 1.422349587101
	fmt.Printf("X to 4 decimal points is %.4f\n", x)
}

func k() {
	x, y, z := 10, 15.5, "Gophers"
	score := []int{10, 20, 30}

	// Print each variable using a specific verb for its type
	fmt.Printf("x is %d, y is %f, z is %s\n", x, y, z)

	// Print the string value enclosed by double quotes ("Gophers")
	fmt.Printf("%q\n", z)

	// Print every variable using the same verb
	fmt.Printf("x is %v, y is %v, z is %v, score is %v\n", x, y, z, score)

	// Print the type of y and score
	fmt.Printf("Y is of type %T and score is of type %T\n", y, score)

}

func j() {
	const (
		Jun = iota + 6
		Jul
		Aug
	)
	fmt.Println(Jun, Jul, Aug)
	//_,_,_ = Jun, Jul, Aug
}

func i() {
	const a int = 7
	const b float64 = 5.6
	const c = float64(a) * b

	x := 8
	_ = x

	// ERRROR ->  You cannot initiate a constant at runtime (constants belong to compile-time)
	//const xc int = x

	// ERRROR ->  You cannot initiate a constant at runtime (constants belong to compile-time)
	//const noIPv6 = math.Pow(2, 128)
}

func h() {
	const x int = 10

	// declaring a constant of type slice int ([]int)
	//const m = []int {1: 3, 4: 5, 6: 8}
	//_ = m
}

func g() {
	const secPerDay = 3600 * 24
	const daysYear = 365
	fmt.Printf("Total number of seconds in a year: %d \n", secPerDay*daysYear)
}

func f() {
	const (
		daysWeek   = 7
		lightSpeed = 299792458
		pi         = 3.14159
	)
}

func e() {
	const daysWeek int = 7
	const lightSpeed int64 = 299792458
	const pi float64 = 3.14159
}

func d() {
	version := "3.1"
	name := "Golang"
	_ = version
	fmt.Println(name)
}

func c() {
	var a float64 = 7.1

	x, y := true, 3.7

	a, x = 5.5, false

	_, _, _ = a, x, y
}

func b() {
	var (
		a int
		b float64
		c bool
		d string
	)

	x, y, z := 20, 15.5, "Gopher"
	_, _, _, _, _, _, _ = a, b, c, d, x, y, z

	//fmt.Println(a, b, c, d)
	//fmt.Println("x is", x, "y is", y, "z is", z)

}

func a() {
	var a int
	var b float64
	var c bool
	var d string

	x := 20
	y := 15.5
	z := "Gopher"

	fmt.Println(a, b, c, d)
	fmt.Println("x is", x, "y is", y, "z is", z)
}
