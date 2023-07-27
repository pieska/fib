package main

import (
	"fib/fib"
	"flag"
	"fmt"
	"os"
)

func main() {

	var n, t, b uint
	var fibs_slice []uint
	var fibs_chan <-chan uint
	var s rune = '['

	flag.UintVar(&n, "n", 5, "Anzahl der Fibonacci-Zahlen")
	flag.UintVar(&t, "t", 0, "Typ 0: Slice, 1: Channel mit range, 2: Channel mit select")
	flag.UintVar(&b, "b", 0, "Channel Buffer Size, 0 ist unbuffered")
	flag.Parse()

	switch t {
	case 0:
		fibs_slice = fib.GenerateSlice(n)

		fmt.Printf("fibs(%d) = %v\n", n, fibs_slice)
	case 1:
		fibs_chan = fib.GenerateChannel(n, b)

		fmt.Printf("fibs(%d) = ", n)

		for fib := range fibs_chan {
			fmt.Printf("%c%d", s, fib)
			s = ' '
		}
		fmt.Println("]")
	case 2:
		fibs_chan = fib.GenerateChannel(n, b)

		fmt.Printf("fibs(%d) = ", n)

		var t uint
		var fib uint

		for isOpen := true; isOpen; {
			select {
			case fib, isOpen = <-fibs_chan:
				if !isOpen {
					break
				}
				fmt.Printf("%c%d", s, fib)
				s = ' '
			default:
				t++
			}
		}
		fmt.Println("]")
		fmt.Printf("waited %d times on channel read\n", t)
	default:
		fmt.Printf("invalid value \"%d\" for flag -t: out of range\n", t)
		flag.Usage()
		os.Exit(1)
	}

	os.Exit(0)
}
