package main

import "fmt"

// Merge bits from two values according to a mask
// https://graphics.stanford.edu/~seander/bithacks.html#MaskedMerge
func main() {
	check(0, 0, 0)
	check(0, 1, 0)
	check(1, 0, 0)
	check(0, 1, 1)
	check(2, 1, 1)
}

func check(a, b, mask int) {
	fmt.Printf("merge(%b, %b, %b) = %b\n", a, b, mask, merge(a, b, mask))
}
func merge(a int, b int, mask int) int {
	return a ^ ((a ^ b) & mask)
}
