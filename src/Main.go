package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(100)
	fmt.Printf("Threads 1: %v\n", runtime.GOMAXPROCS(-1))
}
