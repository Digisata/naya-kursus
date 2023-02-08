package main

import (
	"fmt"
	"math"
)

func isPowerOfTwo(n int) bool {
	square := 0
	res := 0.0

	for int(res) <= n {
		res = math.Pow(2.0, float64(square))
		if int(res) == n {
			return true
		}
		square++
	}

	return false
}

func main() {
	fmt.Println(isPowerOfTwo(16))
}
