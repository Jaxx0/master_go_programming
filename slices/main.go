package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func main() {
	var cities []string
	fmt.Println("citiea is equal to nil:", cities == nil)
	fmt.Printf("Cities %#v\n", cities)
	fmt.Println(len(cities))

	//cities[0] = "London"
	numbers := []int{2, 3, 4, 5, 6, 7}
	fmt.Println(numbers)

	nums := make([]int, 2)
	fmt.Printf("%#v\n", nums)

	type names []string
	friends := names{"Dan", "Maria"}
	fmt.Println(friends)
	fmt.Printf("Friends: %#v\n", friends)

	myFriend := friends[0]
	fmt.Printf("My best friend is %s\n", myFriend)

	friends[0] = "Gabriel"
	fmt.Printf("My best friend is %s\n", friends[0])

	for i, v := range numbers {
		fmt.Printf("index: %v, value: %v\n", i, v)
	}

	var n []int
	n = numbers
	fmt.Println(n)

	fmt.Println("_______________________________________________________________")

	var j []int
	fmt.Println(j == nil)

	k := []int{}
	fmt.Println(k == nil)

	a, b := []int{1, 2, 3}, []int{1, 2, 3}
	//fmt.Println(a == b)

	var eq bool = true

	for i, VakueB := range b {
		if VakueB != a[i] {
			eq = false
			break
		}
	}
	if eq {
		fmt.Println("Slices a and b are equal")
	} else {
		fmt.Println("Slices a and b are NOT equal")
	}

	nambas := []int{2, 3}

	nambas = append(nambas, 10)
	fmt.Println(nambas)

	nambas = append(nambas, 20, 30, 40, 50)
	fmt.Println(nambas)

	c := []int{100, 200}
	nambas = append(nambas, c...)
	fmt.Println(nambas)

	src := []int{10, 20, 30}
	dst := make([]int, 2)
	nn := copy(dst, src)
	fmt.Println("XXXX")
	fmt.Println(src)
	fmt.Println(dst)
	fmt.Println(nn)
	fmt.Println(strings.Repeat("#", 70))
	fmt.Println(src, dst, nn)

	aa := [9]int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	bb := aa[1:3]
	fmt.Printf("%#v, %T\n", bb, bb)
	fmt.Println(aa[2:])
	fmt.Println(aa[2:len(aa)])
	fmt.Println(aa[:3])
	fmt.Println(aa[0:3])
	fmt.Println(aa[:])
	fmt.Println(aa[0:len(aa)])
	s1 := []int{1, 2, 3, 4, 5, 6}
	s1 = append(s1[:4], 100)
	fmt.Println(s1)
	s1 = append(s1[:4], 500)
	fmt.Println(s1)

	// Backing array
	arr1 := [4]int{10, 20, 30, 40}
	slice1, slice2 := arr1[0:2], arr1[1:3]
	fmt.Println(arr1, slice1, slice2)

	arr1[1] = 2
	fmt.Println(arr1, slice1, slice2)

	cars := []string{"Ford", "Honda", "Audi"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)
	fmt.Println(newCars)

	cars[0] = "Nissan"
	fmt.Println(cars, newCars)

	s11 := []int{10, 20, 30, 40, 50}
	newSlice := s11[0:3]
	fmt.Println(len(newSlice), cap(newSlice))

	newSlice = s11[2:5] // {30, 40, 50}
	fmt.Println(len(newSlice), cap(newSlice))

	aaa := [5]int{1, 2, 3, 4, 5}
	bbb := []int{1, 2, 3, 4, 5}
	fmt.Printf("array's size in bytes %d \n", unsafe.Sizeof(aaa))
	fmt.Printf("slice's size in bytes %d \n", unsafe.Sizeof(bbb))

	var men []int
	fmt.Printf("%#v\n", men)
	fmt.Printf("Length: %d, capacity: %d \n", len(men), cap(men))

	men = append(men,1,2)
	fmt.Printf("Length: %d, capacity: %d \n", len(men), cap(men))

	men = append(men,3)
	fmt.Printf("Length: %d, capacity: %d \n", len(men), cap(men))

	men = append(men,4)
	fmt.Printf("Length: %d, capacity: %d \n", len(men), cap(men))

	men = append(men,100)
	fmt.Printf("Length: %d, capacity: %d \n", len(men), cap(men))

	men = append(men,200)
	fmt.Printf("Length: %d, capacity: %d \n", len(men), cap(men))

	men = append(men,500)
	fmt.Printf("Length: %d, capacity: %d \n", len(men), cap(men))

	men = append(men,200)
	fmt.Printf("Length: %d, capacity: %d \n", len(men), cap(men))

	men = append(men,500)
	fmt.Printf("Length: %d, capacity: %d \n", len(men), cap(men))

	letters:= []string{"A", "B", "C", "D", "E", "F"}
	letters = append(letters[0:1], "X", "Y")
	fmt.Println(letters)
	fmt.Println(letters, len(letters), cap(letters))
	//fmt.Println(letters[4])
	fmt.Println(letters[3:6])



}
