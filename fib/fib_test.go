package fib

import (
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
		var fibs <-chan uint = GenerateChannel(5)
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
		GenerateChannel(i)
		i++
	}
}
