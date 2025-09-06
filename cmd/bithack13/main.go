package main

import "fmt"

var bitTable = [256]byte{}

func init() {
	bitTable[0] = 0
	for i := 0; i < 256; i++ {
		bitTable[i] = byte(i&1) + bitTable[i/2]
	}
}

// Counting bits set by lookup table
// https://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetTable
func main() {
	check(1)
	check(15)
}

func check(v uint) {
	fmt.Printf("check '%b' -> '%d'\n", v, countBits(v))
}

func countBits(v uint) byte {
	return bitTable[v&0xff] +
		bitTable[(v>>8)&0xff] +
		bitTable[(v>>16)&0xff] +
		bitTable[v>>24]
}
