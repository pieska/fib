package fib

// Generate ...
func GenerateSlice(count uint) (fibs []uint) {
	var prev, current uint = 0, 1
	var i uint

	for i = 0; i < count; i++ {
		fibs = append(fibs, current)
		prev, current = current, prev+current
	}
	return
}

func GenerateChannel(count uint) (fibs chan (uint)) {

	fibs = make(chan uint)

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
