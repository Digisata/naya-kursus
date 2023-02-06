package main

import "fmt"

func main() {
	numbers := []int{1,-4,7,12}
	fmt.Println(PositiveSum(numbers))
}

func PositiveSum(numbers []int) int {
	result := 0
	for _, number := range numbers {
	  if number > 0 {
		result += number
	  }
	}
	return result
  }