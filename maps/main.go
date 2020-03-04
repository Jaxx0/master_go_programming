package main

import "fmt"

func main()  {
	var employees map[string]string
	fmt.Printf("%#v \n", employees)
	fmt.Printf("Number of pairs: %d \n", len(employees))

	var m2 map[[5]int]string
	fmt.Printf("%#v \n", m2)
	//employees["Dan"] = "Programmer"
	//fmt.Print(employees)

	people := map[string]float64{}
	people["John"] = 21.4
	people["Mary"] = 21.4
	fmt.Println(people)

	map1:=make(map[int]int)
	map1[4]=8
	fmt.Println(map1)

	balances := map[string]float64{
		"USD":564.13,
		"euro": 7667,
		"DED":56,
	}
	fmt.Println(balances)

	balances["USD"] = 900
	balances["GBP"] = 100
	fmt.Println(balances)

	fmt.Println(balances["RON"])

	balances["RON"] = 0

	v, ok := balances["RON"]
	if ok{
		fmt.Println("The RON balance is", v)
	}else{
		fmt.Println("The RON key doesn't exist")
	}

	for k,v:=range balances{
		fmt.Printf("Key: %#v, Value: %v \n", k, v)
	}

	delete(balances, "USD")
	fmt.Println(balances)

}
