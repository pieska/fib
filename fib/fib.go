package fib

// generate Fibonacci and return them as slice
func GenerateSlice(count uint) []uint {
	var prev, current uint = 0, 1
	var i uint

	// create an empty slice with cap = count to prevent enlarging array by copying
	var fibs = make([]uint, 0, count)

	for i = 0; i < count; i++ {
		fibs = append(fibs, current)
		prev, current = current, prev+current
	}
	return fibs
}

// generate Fibonacci and return a receive-only channel to get them
func GenerateChannel(count uint, buffersize uint) <-chan (uint) {

	// create a two-way buffered channel
	var fibs chan uint = make(chan uint, buffersize)

	go func(count uint) {
		var prev, current uint = 0, 1
		var i uint

		defer close(fibs)

		for i = 0; i < count; i++ {
			fibs <- current
			prev, current = current, prev+current
		}
	}(count)

	return fibs
}
