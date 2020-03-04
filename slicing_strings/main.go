package main

import (
	"fmt"
	"strings"
)

func main() {
	// ascii string I byte = 1 character

	s1 := "I love Golang!"
	fmt.Println(s1[2:6])

	// Non ascii
	s2 := "你好"
	fmt.Println(s2[0:2])

	rs := []rune(s2)
	fmt.Printf("%T\n", rs)
	fmt.Println(string(rs))

	//Non-ASCII characters are encoded using more than 1 byte

	result := strings.Contains("I love go programming", "love")
	fmt.Println(result)

	a := strings.ContainsAny("sasa", "abcd")
	fmt.Println(a)

	b := strings.ContainsRune("Golang", 'g')
	fmt.Println(b)

	c := strings.Count("Cheesey", "e")
	fmt.Println(c)

	d := strings.Count("Fiver", "")
	fmt.Println(d)

	fmt.Println(strings.ToLower("GO PYTHON JAVA C"))
	fmt.Println(strings.ToUpper("GO PYTHON JAVA C"))

	fmt.Println("go" == "go")
	fmt.Println("Go" == "go")

	p := fmt.Println
	p(strings.ToLower("Go") == strings.ToLower("go"))

	p(strings.EqualFold("Go", "go"))
	myStr := strings.Repeat("abc", 30)
	p(myStr)

	myStr = strings.Replace("192.168.0.1", ".", ":", -1)
	p(myStr)

	myStr = strings.Replace("Jesus Walks", "Walks", "saves", 2)
	p(myStr)

	myStr = strings.ReplaceAll("ONOMATOPOEIA", "POEIA", "-POW-EY-YEAH")
	p(myStr)

	s := strings.Split("A,B,C,D", ",")
	fmt.Printf("%T\n", s)
	fmt.Printf("%v\n", s)

	s = strings.Split("Go for Go", "")
	fmt.Printf("%v\n", s)
	fmt.Printf("%T\n", s)

	s = []string{"I", "am", "learning", "Golang"}
	myStr = strings.Join(s, "__")
	fmt.Println(myStr)

	myStr = "Oranges Pineapples \n apples Bananas"
	fields := strings.Fields(myStr)
	fmt.Printf("%T\n", fields)
	fmt.Println(fields)

	s22 := strings.TrimSpace("\t Goodbye Windows, Hello Linux \n")
	fmt.Printf("%q\n", s22)

	s21 := strings.Trim("....Hello Gophers!......!@#$%?\\", "!.@#$%\\?")
	fmt.Printf("%q\n", s21)


}
