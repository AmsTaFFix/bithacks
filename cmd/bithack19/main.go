package main

import "fmt"

// Computing parity the naive way
// https://graphics.stanford.edu/~seander/bithacks.html#ParityNaive
func main() {
	check(1)
	check(15)
	check(16)
	check(31)
}

func check(v uint) {
	fmt.Printf("check %v ('%b') -> '%v'\n", v, v, computeParity(v))
}
func computeParity(v uint) bool {
	parity := false

	for v > 0 {
		parity = !parity
		v = v & (v - 1)
	}

	return parity
}
