package main

import "fmt"

// Compute parity of word with a multiply
// https://graphics.stanford.edu/~seander/bithacks.html#%23ParityMultiply
func main() {
	check(1)
	check(15)
	check(16)
	check(31)
	check(255)
	check((1 << 32) + 1)
	check((1 << 32) + 3)
	check((1 << 63))
}

func check(v uint64) {
	if v < 1<<32 {
		fmt.Printf("check uint32(%v) ('%b') -> '%v'\n", v, v, computeParityUint32(uint32(v)))
	} else {
		fmt.Printf("check uint64(%v) ('%b') -> '%v'\n", v, v, computeParityUint64(v))
	}

}

func computeParityUint32(v uint32) uint32 {
	v ^= v >> 1
	v ^= v >> 2
	v = (v & uint32(0x11111111)) * uint32(0x11111111)
	return (v >> 28) & 1
}

func computeParityUint64(v uint64) uint64 {
	v ^= v >> 1
	v ^= v >> 2
	v = (v & uint64(0x1111111111111111)) * uint64(0x1111111111111111)
	return (v >> 60) & 1
}
