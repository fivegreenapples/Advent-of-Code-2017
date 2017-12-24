package main

// This is a Go version of the input assembly code after a few rounds of
// simpllification and working out what it does.
//
// It finds a count of non primes in between two numbers A and B where the only actual
// values considered are A, A+17, A+17+17, A+17+17+17, etc...
func doPart2(initialA int) int {
	a, b, d, f, h := 0, 0, 0, 0, 0
	a = initialA

	var start int
	var stop int

	if a == 0 {
		start = 81
		stop = 81
	} else {
		b = 108100
		start = 108100
		stop = 108100 + 17000
	}
	for b = start; b <= stop; b += 17 {
		f = 1
		for d = 2; d < b; d++ {
			//
			// commented out; this is the main efficiency fail
			//
			// for e = 2; e < b; e++ {
			// 	// "set g d",
			// 	g = d
			// 	// "mul g e",
			// 	g *= e
			// 	numMultiplies++
			// 	// "sub g b",
			// 	g -= b
			// 	// "jnz g 2",
			// 	if g == 0 {
			// 		// "set f 0",
			// 		f = 0
			// 	}
			// 	// f is zero if (d * e) - b == 0
			// 	// --> if is zero if b is not prime
			// }

			// replacement non-prime test
			if b%d == 0 {
				f = 0
				break
			}

		}
		if f == 0 {
			h++
		}
	}

	return h
}
