package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(50)

	for i:=100.; i<150.; i++{
		go func(i float64, wg *sync.WaitGroup) {
			ans := math.Sqrt(i)
			fmt.Printf("Square root of %v is %v \n", i, ans)
			wg.Done()
		}(i, &wg)
	}




	wg.Wait()
}