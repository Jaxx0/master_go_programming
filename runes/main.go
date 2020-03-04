package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main()  {
	var1, var2  := 'a', 'b'
	fmt.Printf("Type: %T, Value: %d \n", var1, var2)

	str:= "tarà"  //country in romanian
	fmt.Println(str)

	fmt.Println(len(str))
	fmt.Println("Byte (not rune) at position 1:", str[1])

	for i:=0; i<len(str); i++{
		fmt.Printf("%c", str[i])
	}

	fmt.Println("\n" + strings.Repeat("#", 100))

	r, size := utf8.DecodeRuneInString(str)
	fmt.Println(r)
	fmt.Println(size)

	fmt.Println("\n" + strings.Repeat("_", 100))
	for _, r := range str{
		fmt.Printf("%c \n", r)
	}

	fmt.Println("\n" + strings.Repeat("_", 100))

	// ASCII

	s11:= "Golang"
	fmt.Println(len(s11))


	// NON ASCII
	name := "Codrutä" // Female name
	fmt.Println(len(name))

	n := utf8.RuneCountInString(name)
	fmt.Println(n)


}
