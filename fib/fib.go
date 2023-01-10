package fib

// Generate ...
func Generate(count uint) (fibs []uint) {
	var prev, current uint = 0, 1
	var i uint

	for i = 0; i < count; i++ {
		fibs = append(fibs, current)
		prev, current = current, prev+current
	}
	return
}
