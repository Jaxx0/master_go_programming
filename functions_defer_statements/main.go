package main

import (
	"fmt"
	"log"
	"os"
)

func foo()  {
	fmt.Println("This is foo")
}

func bar()  {
	fmt.Println("This is bar")
}

func foobar()  {
	fmt.Println("This is foobar")
}

func main()  {
	// Note the LIFO order of deferred functions
	defer foo()
	bar()
	fmt.Println("After deferring foo and printing bar")
	defer foobar()
	file, err := os.Open("main.go")
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()

	//code that works with a file
}