package main

import "fmt"

func main() {
	f3()
}

func f1() {
	var m1 map[string]string
	fmt.Printf("%T\n", m1)
	fmt.Printf("%#v", m1)

	m2 := map[int]string{1: "Kenya", 2: "USSR", 10: "Abba"}
	fmt.Println(m2)

	//r, ok := m2[1]  // Exists
	r, ok := m2[3] // NOT Existent
	if ok {
		fmt.Print("The value with that key is ", r)
	} else {
		fmt.Println("Key does not exist")
	}
}


func f2() {
	//var m1 map[int]bool{}
	//m1[5] = true

	m2 := map[int]int{3: 10, 4: 40}
	m3 := map[int]int{3: 10, 4: 40}

	s2 := fmt.Sprintf("%s", m2)
	s3 := fmt.Sprintf("%s", m3)

	fmt.Println(s2, s3)

	fmt.Println(s2 == s3)
}

func f3()  {
	m := map[int]bool{1: true, 2: false, 3: false}
	delete(m, 3)
	fmt.Println(m)
	for k, v := range m{
		fmt.Printf("Key: %d, Value: %t \n", k,v)
	}
}