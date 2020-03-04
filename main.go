package main



import "fmt"



func main() {

	var y = 458

	var p = &y

	*p = 500

	fmt.Println(y)

}