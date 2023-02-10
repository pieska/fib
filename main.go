package main

import (
	"fib/fib"
	"flag"
	"fmt"
	"os"
)

func main() {

	var n uint

	flag.UintVar(&n, "n", 5, "Stellen der Fibonacci-Zahl")
	flag.Parse()

	fib := fib.Generate(n)

	fmt.Printf("fibs(%d) = %v\n", n, fib)

	os.Exit(0)
}
