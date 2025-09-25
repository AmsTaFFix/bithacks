package main

import (
	"fmt"
)

// BitReverseTable256 will be populated at runtime
var BitReverseTable256 [256]uint8

// init function runs automatically when the package is loaded
func init() {
	generateBitReverseTable()
}

// generateBitReverseTable creates the lookup table by reversing bits for each byte value
func generateBitReverseTable() {
	for i := 0; i < 256; i++ {
		BitReverseTable256[i] = reverseByte(uint8(i))
	}
}

// reverseByte reverses the bits in a single byte
func reverseByte(b uint8) uint8 {
	var result uint8
	for i := 0; i < 8; i++ {
		result = (result << 1) | (b & 1)
		b >>= 1
	}
	return result
}

// Reverse bits in word by lookup table
// https://graphics.stanford.edu/~seander/bithacks.html#BitReverseTable
func main() {
	check(7)
	check(127)
	check(128)
}

func check(v uint) {
	fmt.Printf("reverse(%d == %b) = %b\n", v, v, do(v))
}

func do(v uint) uint {
	c := (uint(BitReverseTable256[v&0xff]) << 24) |
		(uint(BitReverseTable256[(v>>8)&0xff]) << 16) |
		(uint(BitReverseTable256[(v>>16)&0xff]) << 8) |
		uint(BitReverseTable256[(v>>24)&0xff])
	return c
}
