package main

import (
	"fmt"
	"time"
)

func main()  {
	language := "goLang"

	switch language {
	case "python":
		fmt.Println("You are learning python, No need for curly braces use indentation")
	case "Go", "goLang":
		fmt.Println("Good, go Go. Use curly braces")
	case "Java":
		fmt.Println("Java is good too")
	default:
		fmt.Println("Any other programming language is a good start")
	}


	n := 5
	// comparing the result of an expression which is bool to another bool value
	switch true {
	case n%2 == 0:
		fmt.Println("Even!")
	case n%2 != 0:
		fmt.Println("Odd!")
	default:
		fmt.Println("Never here!")
	}

	//** Switch simple statement **//

	// Syntax: statement (n:=10), semicolon and a switch condition
	//(true in this case, we are comparing boolean expressions that return true)
	// we can remove the word "true" because it's the default
	switch n := 10; true {
	case n > 0:
		fmt.Println("Positive")
	case n < 0:
		fmt.Println("Negative")
	default:
		fmt.Println("Zero")
	}


	h := time.Now().Hour()
	fmt.Println(h)

}
