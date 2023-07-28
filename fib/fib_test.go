package fib

import (
	"fmt"
	"reflect"
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

	var i uint = 0

	for i < uint(b.N) {
		GenerateSlice(i)
		i++
	}
}

func BenchmarkGenerateChannel(b *testing.B) {

	var i uint = 0

	for i < uint(b.N) {
		GenerateChannel(i, 0)
		i++
	}
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
