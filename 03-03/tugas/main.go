package main

import "fmt"

func plusOne(digits []int) (res []int) {
	carry := 0
	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + carry
		if i == len(digits)-1 {
			sum++
		}
		carry = sum / 10
		res = append([]int{sum % 10}, res...)
	}
	if carry != 0 {
		res = append([]int{carry}, res...)
	}
	return
}

func main() {
	fmt.Println(plusOne([]int{1,2,3}))
}
