package main

import (
	"fmt"
	"sort"
)

func IsTriangle(a, b, c int) bool {
    sides := []int{a, b, c}
    sort.Ints(sides)
    return sides[0] + sides[1] > sides[2]
}

func main() {
	fmt.Println(IsTriangle(5, 1, 2))
	fmt.Println(IsTriangle(4, 2, 3))
}