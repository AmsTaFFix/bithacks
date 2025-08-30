package main

import (
	"fmt"
	"strconv"
)

const charBit = 8

// Compute the integer absolute value (abs) without branching
// https://graphics.stanford.edu/~seander/bithacks.html#IntegerAbs
func main() {
	abs(1)
	abs(0)
	abs(-1)
	abs(-90000000000)
}

func abs(v int) {
	mask := v >> (strconv.IntSize*charBit - 1)

	normal := uint((v + mask) ^ mask)
	patented := uint((v ^ mask) - mask)

	fmt.Printf("abs(%d) = %d and %d;\n", v, normal, patented)
}
