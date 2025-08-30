package main

import "fmt"

// Determining if an integer is a power of 2
// https://graphics.stanford.edu/~seander/bithacks.html#DetermineIfPowerOf2
func main() {
	check(-1)
	check(0)
	check(1)
	check(2)
	check(-2)
	check(-4)
	check(8)
	check(9)
	check(10)

}

func check(v int) {
	fmt.Printf("%d: %v / %v\n", v, (v&(v-1)) == 0, v > 0 && (v&(v-1)) == 0)
}
