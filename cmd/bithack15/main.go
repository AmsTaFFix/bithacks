package main

import "fmt"

// Counting bits set in 14, 24, or 32-bit words using 64-bit instructions
// https://graphics.stanford.edu/~seander/bithacks.html#CountBitsSet64
func main() {
	check(1)
	check(15)
}

func check(v uint) {
	fmt.Printf("check '%b' -> '%d'\n", v, countBits(uint64(v)))
}

func countBits(v uint64) uint64 {
	c := uint64(0)
	// option 1, for at most 14-bit values in v:
	c = (v * uint64(0x200040008001) & uint64(0x111111111111111)) % 0xf

	// option 2, for at most 24-bit values in v:
	c = ((v & 0xfff) * uint64(0x1001001001001) & uint64(0x84210842108421)) % 0x1f
	c += (((v & 0xfff000) >> 12) * uint64(0x1001001001001) & uint64(0x84210842108421)) % 0x1f

	// option 3, for at most 32-bit values in v:
	c = ((v & 0xfff) * uint64(0x1001001001001) & uint64(0x84210842108421)) % 0x1f
	c += (((v & 0xfff000) >> 12) * uint64(0x1001001001001) & uint64(0x84210842108421)) %
		0x1f
	c += ((v >> 24) * uint64(0x1001001001001) & uint64(0x84210842108421)) % 0x1f

	return c
}
