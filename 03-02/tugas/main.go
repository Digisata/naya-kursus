package main

import (
	"fmt"
)

func GetSum(a, b int) (sum int) {
	if a == b {
		return a
	}

	if a > b {
		a, b = b, a
	}

	for i := a; i <= b; i++ {
		sum += i
	}
	return
}

func main() {
	fmt.Println(GetSum(505, 4))
}
