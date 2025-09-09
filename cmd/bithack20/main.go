package main

import "fmt"

// ParityTable256 is a lookup table for computing parity of a byte
var ParityTable256 [256]bool

func init() {
	// Simple approach: compute parity for each byte value
	for i := 0; i < 256; i++ {
		ParityTable256[i] = computeParityByte(uint8(i))
	}
}

// Helper function to compute parity of a single byte
func computeParityByte(v uint8) bool {
	parity := false
	for v > 0 {
		parity = !parity
		v = v & (v - 1) // Clear the lowest set bit
	}
	return parity
}

// Compute parity by lookup table
// https://graphics.stanford.edu/~seander/bithacks.html#ParityLookupTable
func main() {
	check(1)
	check(15)
	check(16)
	check(31)
	check(256)
}

func check(v uint) {
	if v < 256 {
		fmt.Printf("check uint8(%v) ('%b') -> '%v'\n", v, v, computeParityUInt8(uint8(v)))
	}
	fmt.Printf("check uint(%v) ('%b') -> '%v'\n", v, v, computeParityUInt(v))
}

func computeParityUInt8(v uint8) bool {
	return ParityTable256[v]
}

func computeParityUInt(v uint) bool {
	v ^= v >> 16
	v ^= v >> 8
	return ParityTable256[v&0xff]
}
