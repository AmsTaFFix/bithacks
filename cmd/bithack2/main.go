package main

import "fmt"

// Detect if two integers have opposite signs
// https://graphics.stanford.edu/~seander/bithacks.html#DetectOppositeSigns
// int x, y;               // input values to compare signs
//
// bool f = ((x ^ y) < 0); // true iff x and y have opposite signs
func main() {
	checkOppositeSign(1, 1)
	checkOppositeSign(1, 0)
	checkOppositeSign(0, 0)
	checkOppositeSign(-1, -1)
	checkOppositeSign(-1, 0)
	checkOppositeSign(-1, 1)
}

func checkOppositeSign(a, b int) {
	fmt.Printf("a=%d, b=%d: Is opposite sign=%v\n", a, b, a^b < 0)
}
