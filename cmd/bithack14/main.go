package main

import "fmt"

// Counting bits set, Brian Kernighan's way
// https://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetKernighan
func main() {
	check(1)
	check(15)
}

func check(v uint) {
	fmt.Printf("check '%b' -> '%d'\n", v, countBits(v))
}

func countBits(v uint) uint {
	c := uint(0)
	for c = 0; v != 0; c++ {
		v &= v - 1
	}

	return c
}
