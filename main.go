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
	var fibs_slice []uint
	var fibs_chan <-chan uint

	flag.UintVar(&n, "n", 5, "Anzahl der Fibonacci-Zahlen")
	flag.BoolVar(&i, "i", false, "immediate output, verwendet Channel")
	flag.Parse()

	if !i {
		fibs_slice = fib.GenerateSlice(n)

		fmt.Printf("fibs(%d) = %v\n", n, fibs_slice)
	} else {
		fibs_chan = fib.GenerateChannel(n)

		fmt.Printf("fibs(%d) =  ", n)

		for fib := range fibs_chan {
			fmt.Printf("%d ", fib)
		}
		fmt.Println()
	}

	os.Exit(0)
}
