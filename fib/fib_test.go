package fib

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

func TestGenerateSlice(t *testing.T) {
	var want []uint = []uint{1, 1, 2, 3, 5}
	var got []uint = GenerateSlice(5)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func TestGenerateChannel(t *testing.T) {
	var want []uint = []uint{1, 1, 2, 3, 5}
	var got []uint = func() (x []uint) {
		var fibs <-chan uint = GenerateChannel(5, 0)
		for fib := range fibs {
			x = append(x, fib)
		}
		return x
	}()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func BenchmarkGenerateSlice(b *testing.B) {

	var i uint
	var r []uint

	sizes := []uint{10, 100, 1000, 10000, 100000}

	for _, s := range sizes {

		b.Run(fmt.Sprintf("GenerateSlice(%d)", s), func(b *testing.B) {
			for i = 0; i < uint(b.N); i++ {
				r = GenerateSlice(uint(s))
			}
			runtime.KeepAlive(r)
		})
	}
}

func BenchmarkGenerateChannel(b *testing.B) {

	var i uint
	var r <-chan uint

	for i = 0; i < uint(b.N); i++ {
		r = GenerateChannel(i, 0)
	}
	runtime.KeepAlive(r)
}

func ExampleGenerateSlice() {
	r := GenerateSlice(5)
	fmt.Println(r)
	// Output:
	// [1 1 2 3 5]
}

func ExampleGenerateChannel() {
	rs := GenerateChannel(5, 0)
	for r := range rs {
		fmt.Println(r)
		// Output:
		// 1
		// 1
		// 2
		// 3
		// 5
	}
}
