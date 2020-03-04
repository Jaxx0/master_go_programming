package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Main execution started")
	fmt.Println("No of CPUS: ", runtime.NumCPU())
	fmt.Println("No of Goroutines: ", runtime.NumGoroutine())
	fmt.Println("OS: ", runtime.GOOS)
	fmt.Println("Arch: ", runtime.GOARCH)
	fmt.Println("GOMAXPROCS: ", runtime.GOMAXPROCS(0))
}
