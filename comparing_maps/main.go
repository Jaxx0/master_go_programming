package main

import "fmt"

func main()  {
	a:=map[string]string{"A":"x"}
	b:=map[string]string{"A":"x"}

	//fmt.Println(a == b)

	s1:=fmt.Sprintf("%s", a)
	s2:=fmt.Sprintf("%s", b)
	fmt.Println(s1, s2)
	fmt.Println(s1 == s2)


	friends := map[string]int{"Dan":25, "Maria":40}
	neighbours := friends

	friends["Dan"] = 50
	fmt.Println(neighbours)

	// copy a map
	people := make(map[string]int)
	for k,v := range friends{
		people[k] =v
	}

		// verification
	people["Anne"] = 18
	fmt.Printf("%#v ___________%#v\n", people, friends)

	delete(friends, "Dan")
	fmt.Println(friends)

	delete(people, "Andre")

}
