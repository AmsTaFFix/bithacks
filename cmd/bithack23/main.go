package main

import "fmt"

// Compute parity in parallel
// https://graphics.stanford.edu/~seander/bithacks.html#ParityParallel
func main() {
	check(1)
	check(15)
	check(16)
	check(31)
	check(255)
	check(1 << 31)
	check((1 << 31) + 1)
}

func check(v uint32) {
	fmt.Printf("check uint32(%v) ('%b') -> '%v'\n", v, v, computeParityUint32(v))

}

func computeParityUint32(v uint32) uint32 {
	v ^= v >> 16
	v ^= v >> 8
	v ^= v >> 4
	v &= 0xf
	return (0x6996 >> v) & 1
}
