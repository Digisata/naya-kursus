package main

import "fmt"

func main() {
	fmt.Println(MakeNegative(5))
	fmt.Println(MakeNegative(-10))
	fmt.Println(MakeNegative(18))
}

func MakeNegative(x int) int {
	if (x < 0) {
	  return x
	}
	return -x
  }