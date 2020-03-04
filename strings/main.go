package main

import "fmt"

func main()  {
	s1 := "Hello there !"
	fmt.Println(s1)

	fmt.Println("He says: \"Hello\"")
	fmt.Println(`He says: "Hello"`)

	s2 := `I like \n Go`
	fmt.Println(s2)

	fmt.Println(`C:\Users\jackson.kitili`)
	fmt.Println("C:\\Users\\jackson.kitili")

	var s3 = "I love " + "Go " + "Programming"
	fmt.Println(s3 + "!")

	fmt.Println("Element at index 0:", s3[0])
	//s3[5] = "x"

	fmt.Printf("%s\n", s3)
	fmt.Printf("%q\n", s3)


}
