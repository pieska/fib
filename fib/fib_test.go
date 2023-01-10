package fib

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	var want []uint = []uint{1, 1, 2, 3, 5}
	var got []uint = Generate(5)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

func BenchmarkGenerate(b *testing.B) {

	const iterations uint = 1_000
	var i uint = 0

	for i < iterations {
		Generate(i)
		i++
	}
}
