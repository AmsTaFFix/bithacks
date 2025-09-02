package main

import "fmt"

// Conditionally set or clear bits without branching
// https://graphics.stanford.edu/~seander/bithacks.html#ConditionalSetOrClearBitsWithoutBranching
func main() {
	check(0, 1, 3)
	check(1, 1, 2)
}

func check(isClear int, m uint, w uint) {
	fmt.Printf("isClear=%d, m=%d, w=%d; f1=%d; f2=%d\n", isClear, m, w, func1(isClear, m, w), func2(isClear, m, w))
}

func func1(isClear int, m uint, w uint) uint { // conditional flag
	w ^= uint(-isClear^int(w)) & m

	return w
}

func func2(isClear int, m uint, w uint) uint { // conditional flag
	w = (w & ^m) | uint(-isClear&int(m))

	return w
}

// // OR, for superscalar CPUs:
//

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}
