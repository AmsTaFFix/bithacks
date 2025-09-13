package main

import "fmt"

// Swapping individual bits with XOR
// https://graphics.stanford.edu/~seander/bithacks.html#SwappingBitsXOR
func main() {
	check(1, 10, 4, 123823489123901)
}

func check(i, j, n, b uint) {
	fmt.Printf("swap(%d, %d, %d, %b) = %b\n", i, j, n, b, swap(i, j, n, b))

}

func swap(i, j, n, b uint) uint {
	//var i, j uint // positions of bit sequences to swap
	//var n uint    // number of consecutive bits in each sequence
	//var b uint    // bits to swap reside in b
	var r uint // bit-swapped result goes here

	var x uint = ((b >> i) ^ (b >> j)) & ((uint(1) << n) - 1) // XOR temporary
	r = b ^ ((x << i) | (x << j))

	return r
}
