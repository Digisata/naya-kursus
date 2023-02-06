package main

import "fmt"

func main() {
	fmt.Println(IsDivisible(20, 5, 2))
	fmt.Println(IsDivisible(25, 5, 7))
	fmt.Println(IsDivisible(100, 20, 2))
}

func IsDivisible(n, x, y int) bool {
    return n % x == 0 && n % y == 0
}