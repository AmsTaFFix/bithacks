package main

import "fmt"

// Compute the minimum (min) or maximum (max) of two integers without branching
// https://graphics.stanford.edu/~seander/bithacks.html#IntegerMinOrMax
func main() {
	check(0, 0)
	check(-1, 0)
	check(0, -1)
	check(0, 1)
	check(1, 1)
	check(-1, 1)
}

func check(x, y int) {
	fmt.Printf("x = %d, y = %d\n", x, y)
	fmt.Printf("min=%d, max=%d\n", min2(x, y), max2(x, y))
}

func min2(x, y int) int {
	return y ^ ((x ^ y) & -bool2int1(x < y)) // min(x, y)
}

func max2(x, y int) int {
	return x ^ ((x ^ y) & -bool2int1(x < y))
}

func bool2int1(b bool) int {
	if b {
		return 1
	}
	return 0
}

func bool2int2(b bool) int {
	// The compiler currently only optimizes this form.
	// See issue 6011.
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	return i
}
