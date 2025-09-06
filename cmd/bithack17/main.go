package main

import "fmt"

// Count bits set (rank) from the most-significant bit upto a given position
// https://graphics.stanford.edu/~seander/bithacks.html#%23CountBitsFromMSBToPos
func main() {
	check(1)
	check(15)
}

func check(v uint) {
	fmt.Printf("check '%b' -> '%d'\n", v, countBits(v))
}
func countBits(v uint) uint {
	// this thing is too much for me for now...
	//   uint64_t v;       // Compute the rank (bits set) in v from the MSB to pos.
	//  unsigned int pos; // Bit position to count bits upto.
	//  uint64_t r;       // Resulting rank of bit at pos goes here.
	//
	//  // Shift out bits after given position.
	//  r = v >> (sizeof(v) * CHAR_BIT - pos);
	//  // Count set bits in parallel.
	//  // r = (r & 0x5555...) + ((r >> 1) & 0x5555...);
	//  r = r - ((r >> 1) & ~0UL/3);
	//  // r = (r & 0x3333...) + ((r >> 2) & 0x3333...);
	//  r = (r & ~0UL/5) + ((r >> 2) & ~0UL/5);
	//  // r = (r & 0x0f0f...) + ((r >> 4) & 0x0f0f...);
	//  r = (r + (r >> 4)) & ~0UL/17;
	//  // r = r % 255;
	//  r = (r * (~0UL/255)) >> ((sizeof(v) - 1) * CHAR_BIT);

	return 0
}
