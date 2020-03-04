package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	f6()
}


func f1() {
	var name string
	name = "Jackson"

	country := "Kenya"
	fmt.Println("Your name: ", name)
	fmt.Println("Country: ", country)

	fmt.Println(`He says: "Hello"`)
	fmt.Println("He says: \"Hello\"")
	fmt.Println(" C:\\Users\\a.txt")
	fmt.Printf("Your name: %s\nCountry: %s\n", name, country)
}

func f2()  {
	s := "ă"
	r := []rune(s)
	fmt.Printf("%T\n", r)

	s1 := "ma"
	s2 := "m"

	fin := s1 + s2 + string(r)
	final := s1 + s2 + s
	fmt.Println(fin)
	fmt.Println(final)
}

func f3()  {
	s1 := "țară means country in Romanian"
	fmt.Println(len(s1))

	// iterating over the string and print out byte by byte
	fmt.Printf("Bytes in string: ")
	for i:=0; i<len(s1); i++{
		fmt.Printf("%v", s1[i])
	}

	fmt.Println()

	//fmt.Printf("Bytes in string: ")
	//for _,v := range s1{
	//	fmt.Printf("%c", v)
	//}

	fmt.Println()

	// iterating over the string and print out rune by rune
	// and the byte index where the rune starts in the string
	for i, r := range s1 {
		fmt.Printf("byte index: %d -> rune: %c\n", i, r)
	}
}

func f4() {
	s1 := "Go is cool!"
	x := s1[0]
	fmt.Println(x)
	fmt.Printf("%c \n", x)

	//s1[0] = 'x'

	// printing the number of runes in the string
	//fmt.Println(len(s1))
	fmt.Println(utf8.RuneCountInString(s1))
}

func f5()  {
	s := "你好 Go!"
	b := []byte(s)
	fmt.Printf("%T \n", b)
	// printing out the byte slice
	fmt.Printf("%#v\n", b)

	// iterating over the byte slice and printing out each index and byte in the byte slice
	for i, v := range b {
		fmt.Printf("%d -> %d\n", i, v)
	}

	for i:=0; i < len(b); i++{
		fmt.Printf("Index: %d -> Byte: %v \n", i, b[i])
	}
}

func f6()  {
	s := "你好 Go!"
	r := []rune(s)
	fmt.Printf("%T \n", r)
	fmt.Printf("%#v \n", r)


	// iterating over the rune slice and printing out each index and rune in the rune slice
	for i, v := range r {
		fmt.Printf("%d -> %c\n", i, v)
	}

	for i:=0; i<len(r); i++{
		fmt.Printf("Index: %d, Byte: %c \n", i, r[i])
	}
}