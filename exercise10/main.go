package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func cube(a float64)float64 {
	return math.Pow(a, 3)
}

func factorial(n uint) uint {
	if n <= 1{
		return 1
	}
	return n * factorial(n - 1)
}


func f1(n uint) (uint, uint) {
	sum := uint(0)
	for i:=uint(0); i<= n; i++{
		sum += i
	}
	f := factorial(n)
	return f, sum
}

func f3(n uint) (int, int)  {
	fact := 1
	sum := 0
	for i:=1; i<=int(n); i++{
		sum += i
		fact *= i
	}
	return fact, sum
}

func f4(n string) int {
	n1, err := strconv.Atoi(n)
	if err != nil{
		log.Fatal()
	}
	n2, _ := strconv.Atoi(n+n)
	n3, _ := strconv.Atoi(n+n+n)
	return n1+n2+n3
}

func variadicSum(a...int) int {
	sum := 0
	for _, v := range a{
		sum += v
	}
	return sum
}

func nakedVariadicSum(a...int) (s int) {
	for _, v := range a{
		s += v
	}
	return
}

func searchItem(a []string, b string) bool {
	for _,v := range a{
		if strings.ToLower(b) != strings.ToLower(v){
			continue
		} else{
			return true
		}
	}
	return false
}

func sensitiveSearchItem(a []string, b string) bool {
	for _,v := range a{
		if strings.EqualFold(v, b){
			return true
		}
	}
	return false
}

func print(msg string) {
	fmt.Println(msg)
}

func copyCat(n int) {
	fmt.Println(n)
}

func main() {
	ans := cube(2.)
	fmt.Println(ans)

	// using recursion
	//f, s := f1(4)
	//fmt.Println(f, s)

	f, s := f3(4)
	fmt.Println(f, s)

	sum := f4("5")
	fmt.Println(sum)

	vSum := variadicSum(1,2,3,4,5,6,7,8,9)
	fmt.Println(vSum)

	nS := nakedVariadicSum(1,2,3,4,5,6,7,8,9)
	fmt.Println(nS)

	animals := []string{"lion", "tiger", "bear"}
	result := searchItem(animals, "bear")
	fmt.Println(result) // => true
	result = searchItem(animals, "pig")
	fmt.Println(result)

	animals = []string{"lion", "tiger", "bear"}
	result = sensitiveSearchItem(animals, "bear")
	fmt.Println(result) // => true
	result = sensitiveSearchItem(animals, "PIG")
	fmt.Println(result)
	result = sensitiveSearchItem(animals, "BEar")
	fmt.Println(result)

	//defer print("The Go gopher is the iconic mascot of the Go project.")
	fmt.Println("Hello, Go playground!")

	echo := copyCat
	fmt.Printf("%T\n", echo)
	echo(5667)

}

