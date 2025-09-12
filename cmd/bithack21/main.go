package main

import "fmt"

// Compute parity of a byte using 64-bit multiply and modulus division
// https://graphics.stanford.edu/~seander/bithacks.html#ParityWith64Bits
func main() {
	check(1)
	check(15)
	check(16)
	check(31)
	check(255)
}

func check(v uint8) {
	fmt.Printf("check uint8(%v) ('%b') -> '%v'\n", v, v, computeParityUInt8(v))
}

func computeParityUInt8(v uint8) uint64 {
	return (((uint64(v) * uint64(0x0101010101010101)) & uint64(0x8040201008040201)) % 0x1FF) & 1
}
