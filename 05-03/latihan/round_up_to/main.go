package main

import "fmt"

func RoundToNext5(n int) int {
	for {
		if n%5 == 0 {
			return n
		}
		n++
	}
}

func main() {
	fmt.Println(RoundToNext5(12))
}
