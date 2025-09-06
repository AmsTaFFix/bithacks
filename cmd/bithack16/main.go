package main

import "fmt"

// Counting bits set, in parallel
// https://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetParallel
func main() {
	check(1)
	check(15)
}

func check(v uint) {
	fmt.Printf("check '%b' -> '%d'\n", v, countBits(v))
}

var S = []uint{1, 2, 4, 8, 16} // Magic Binary Numbers
var B = []uint{0x55555555, 0x33333333, 0x0F0F0F0F, 0x00FF00FF, 0x0000FFFF}

func countBits(v uint) uint {
	c := uint(0)
	c = v - ((v >> 1) & B[0])
	c = ((c >> S[1]) & B[1]) + (c & B[1])
	c = ((c >> S[2]) + c) & B[2]
	c = ((c >> S[3]) + c) & B[3]
	c = ((c >> S[4]) + c) & B[4]

	return c
}
