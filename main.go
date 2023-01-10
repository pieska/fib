package main

import (
	"fib/fib"
	"fmt"
	"os"
)

func main() {
	const n uint = 5

	var fib []uint = fib.Generate(n)

	fmt.Printf("fibs(%d) = %v\n", n, fib)

	os.Exit(0)
}
