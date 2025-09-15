package main

import (
	"fmt"
	"unsafe"
)

// Reverse bits the obvious way
// https://graphics.stanford.edu/~seander/bithacks.html#BitReverseObvious
func main() {
	check(7)
	check(127)
	check(128)
}

func check(v uint) {
	fmt.Printf("reverse(%d == %b) = %b\n", v, v, do(v))

}

func do(v uint) uint {
	r := v
	s := unsafe.Sizeof(v)*8 - 1 // extra shift needed at end

	for v >>= 1; v > 0; v >>= 1 {
		r <<= 1
		r |= v & 1
		s--
	}
	r <<= s // shift when v's highest bits are zero

	return r
}
