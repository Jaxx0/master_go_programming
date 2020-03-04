package main

import "fmt"

type vehicle interface {
	License() string
	Name() string
}

type car struct {
	licenceNo string
	brand     string
}

func (c car) License() string  {
	return c.licenceNo
}

func (c car) Name() string  {
	return c.brand
}

func main() {
	var myCar vehicle
	myCar = car{licenceNo:"KCX 723 X", brand:"Audi"}
	a := myCar.License()
	b := myCar.Name()
	fmt.Println(a)
	fmt.Println(b)

}