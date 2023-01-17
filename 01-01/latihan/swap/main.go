package main

import (
	"fmt"
)

func main() {
	var x = 20
	var y = 30

	fmt.Println("x =", x)
	fmt.Println("y =", y)

	fmt.Println("Swap")

	x, y = y, x

	fmt.Println("x =", x)
	fmt.Println("y =", y)
}
