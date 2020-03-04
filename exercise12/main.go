package main

import "fmt"

type money float64

func (m money) print() {
	fmt.Printf("%.2f\n", m)
}
func (m money) printStr() string {
	return fmt.Sprintf("%.2f\n", m)
}

type book struct {
	price float64
	title string
}

func (b book) vat() float64 {
	return b.price * 0.09
}

func (b *book) discount() float64  {
	return b.price * 0.9
}

func main() {
	var salary money = 9 *5.989890
	fmt.Println(salary)
	salary.print()
	fmt.Printf("Type: %T\n", salary)

	s:= salary.printStr()
	fmt.Println(s)
	fmt.Printf("Type: %T\n", s)

	tis := book{price:100, title:"'Tis - Frank McCourt"}
	fmt.Println(tis.vat())
	fmt.Println(tis.discount())
}

