package main

import (
	"fmt"
	"sort"
)

func Gimme(array [3]int) (res int) {
	sortedArray := array
	sort.Ints(sortedArray[:])
	for i, val := range array {
		if val == sortedArray[1] {
			res = i
			break
		}
	}
	return
}

func main() {
	fmt.Println(Gimme([3]int{2, 3, 1}))
}
