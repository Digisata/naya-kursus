package main

import "fmt"

func century(year int) int {
	if year%100 == 0 {
		return year / 100
	}
	return int(year/100) + 1
}

func main() {
	fmt.Println(century(1990))
}
