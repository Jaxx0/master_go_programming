package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	n := 144.

	var wg sync.WaitGroup
	wg.Add(1)

	go func(n float64, wg *sync.WaitGroup) {
		ans := math.Pow(n, 0.5)
		ans = math.Sqrt(n)
		fmt.Printf("Square root of %v is %v \n", n, ans)
		wg.Done()
	}(n, &wg)
	wg.Wait()
}