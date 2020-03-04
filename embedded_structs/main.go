package main

import (
	"fmt"
	"strings"
)

func main()  {
	type Contact struct {
		email, address string
		phone int
	}

	type Employee struct {
		name string
		salary int
		contactInfo Contact
	}

	john := Employee{
		name:        "Jon Snow",
		salary:      900000,
		contactInfo: Contact{
			email:"jonsnow@got.ent",
			address:"32 Winterfell",
			phone: 98759835,
		},
	}

	fmt.Printf("%#v\n", john)
	fmt.Println(strings.Repeat("#", 50))
	fmt.Println(john)
	fmt.Printf("Employees email: %s\n", john.contactInfo.email)

	john.contactInfo.email = "new_email@gmail.com"
	fmt.Printf("Employees new email: %s\n", john.contactInfo.email)

	myContact := Contact{"andrei@gmail.com", "45 Boulevard Drive Romania", 9876868 }
	fmt.Println(myContact)

}
