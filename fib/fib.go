package fib

// Generate ...
func GenerateSlice(count uint) []uint {
	var prev, current uint = 0, 1
	var i uint

	// create an empty slice with cap = count
	var fibs = make([]uint, 0, count);

	for i = 0; i < count; i++ {
		fibs = append(fibs, current)
		prev, current = current, prev+current
	}
	return fibs
}

// return a receive-only channel only
func GenerateChannel(count uint) <-chan (uint) {

	// create a two-way channel
	var fibs chan uint = make(chan uint)

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
