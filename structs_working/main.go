package main

import "fmt"

func main()  {
	type book struct {
		title string
		author string
		year int
	}

	lastBook := book{title:"Anna Karerina"}
	fmt.Println(lastBook.title)

	//page := lastBook.pages
	fmt.Printf("%#v\n", lastBook)
	lastBook.author = "Leo Tolstoy"
	lastBook.year = 1778

	fmt.Printf("%#v\n", lastBook)

	aBook := book{title:"Anna Karerina", author:"Leo Tolstoy", year:1778}
	fmt.Println(aBook == lastBook)

	// copy of a struct
	myBook := aBook
	myBook.year = 2020
	fmt.Println(myBook, aBook)


	//Anonymous struct types
	diana := struct {
		firstName, lastName string
		age int
	}{firstName:"Diana",
		lastName:"Muller",
		age:30}

	fmt.Printf("%#v\n", diana)
	fmt.Printf(" Diana's age: %d\n", diana.age)

	type Book struct {
		string
		float64
		bool
	}

	b1 := Book{"1984 by George", 10.32, false}
	fmt.Printf("%#v\n", b1)
	fmt.Println(b1.string)

	type Employee struct {
		name string
		salary int
		bool
	}

	e:= Employee{"John", 8000000, true}
	fmt.Printf("%#v\n", e)

}

