package main

import "fmt"

// Conditionally negate a value without branching
// https://graphics.stanford.edu/~seander/bithacks.html#ConditionalNegate
func main() {
	check(10)
	check(-10)
	check(0)
}

func check(v int) {
	fmt.Printf("check '%d'\n", v)
	fmt.Printf("shouldNegateFlag: false: %v; true: %v\n", flagShouldNegate(v, 0),
		flagShouldNegate(v, 1))
	fmt.Printf("shouldNotNegateFlag: false: %v; true: %v\n", flagShouldNotNegate(v, 0),
		flagShouldNotNegate(v, 1))
}

func flagShouldNegate(v int, fNegate int) int {
	return (v ^ -fNegate) + fNegate
}
func flagShouldNotNegate(v int, fDontNegate int) int {
	return (fDontNegate ^ (fDontNegate - 1)) * v
}
