package main

import "fmt"

func main() {
	f4()
}

func f1() {
	type person struct {
		name string
		age int
		colours []string
	}

	me := person{name:"Michael Gray", age:34, colours: []string{"blue", "red"}}
	you := person{name:"Thomas Shelby Gray", age:43, colours: []string{"Gray", "red"}}

	fmt.Printf("%v\n", me)
	fmt.Printf("%+v\n", you)
}

func f2()  {
	type person struct {
		name string
		age int
		colours []string
	}

	me := person{name:"Michael Gray", age:34, colours: []string{"blue", "red"}}
	you := person{name:"Thomas Shelby Gray", age:43, colours: []string{"Gray", "red"}}
	me.name = "Andrei"

	fmt.Printf("%+#v \n", you.colours)
	you.colours = append(you.colours, "maroon")
	fmt.Printf("%+v\n", me)
	fmt.Printf("%+v\n", you)

	var colors []string = you.colours
	fmt.Printf("Type: %T, Value: %v\n", colors, colors)

}

func f3()  {
	type person struct {
		name string
		age int
		colours []string
	}

	me := person{name:"Michael Gray", age:34, colours: []string{"blue", "red"}}
	you := person{name:"Thomas Shelby Gray", age:43, colours: []string{"Gray", "red"}}
	_ = you

	for _, c := range me.colours{
		fmt.Println(c)
	}

}

func f4()  {
	type grades struct {
		grade string
		course string
	}

	type person struct {
		name string
		age int
		colours []string
		grades grades
	}


	me := person{
		name:"Michael Gray",
		age:34,
		colours: []string{"blue", "red"},
		grades:grades{grade:"A", course:"Engineering"}}

	you := person{
		name:"Thomas Shelby Gray",
		age:43,
		colours: []string{"Gray", "red"},
		grades:grades{grade:"B", course:"Sociology"}}

	me.grades.course = "Golang"
	me.grades.grade = "98"

	fmt.Printf("%+v\n", me)
	fmt.Printf("%+v\n", you)

}