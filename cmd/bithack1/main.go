package main

import (
	"fmt"
	"unsafe"
)

// Compute the sign of an integer
// https://graphics.stanford.edu/~seander/bithacks.html#CopyIntegerSign
func main() {
	v := -23483
	sign := 0
	charBit := 8
	sign = -int(uint(v) >> (int(unsafe.Sizeof(v))*charBit - 1))
	fmt.Printf("%d\n", sign)

	sign = v >> (int(unsafe.Sizeof(v))*charBit - 1)
	fmt.Printf("%d\n", sign)
}

//int v;      // we want to find the sign of v
//int sign;   // the result goes here

// CHAR_BIT is the number of bits per byte (normally 8).
//sign = -(v < 0);  // if v < 0 then -1, else 0.
// or, to avoid branching on CPUs with flag registers (IA32):
//sign = -(int)((unsigned int)((int)v) >> (sizeof(int) * CHAR_BIT - 1));
// or, for one less instruction (but not portable):
//sign = v >> (sizeof(int) * CHAR_BIT - 1);
