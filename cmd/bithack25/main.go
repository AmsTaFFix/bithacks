package main

import "fmt"

// Swapping values with XOR
// https://graphics.stanford.edu/~seander/bithacks.html#SwappingValuesXOR
func main() {
	check(0, 0)
	check(2, 1)
	check(19, 19043)
}

func check(a, b int) {
	fmt.Printf("swap(%d, %d)", a, b)
	swap(&a, &b)
	fmt.Printf(" -> %d, %d\n", a, b)

}

func swap(a, b *int) {
	if a == b {
		return
	}
	*a ^= *b
	*b ^= *a
	*a ^= *b
}
