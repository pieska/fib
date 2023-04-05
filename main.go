package main

import (
	"fib/fib"
	"flag"
	"fmt"
	"os"
)

func main() {

	var n uint
	var i bool

	flag.UintVar(&n, "n", 5, "Anzahl der Fibonacci-Zahlen")
	flag.BoolVar(&i, "i", false, "immediate output, verwendet Channel")
	flag.Parse()

	if !i {
		fibs := fib.GenerateSlice(n)

		fmt.Printf("fibs(%d) = %v\n", n, fibs)
	} else {
		fibs := fib.GenerateChannel(n)

		fmt.Printf("fibs(%d) =  ", n)

		for fib := range fibs {
			fmt.Printf("%d ", fib)
		}
		fmt.Println()
	}

	os.Exit(0)
}
