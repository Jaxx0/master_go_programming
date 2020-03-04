package main

import (
	"fmt"
	"sync"
)

func sum(a, b float64, wg *sync.WaitGroup) {
	ans := a+b
	fmt.Printf("Sum is : %.2f \n", ans)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go sum(3.142, 14.562, &wg)
	go sum(3.56, 78857.59, &wg)
	go sum(993.56, 267.875889, &wg)
	wg.Wait()

}