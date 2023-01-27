package main

import "fmt"

func getMinLength(a []int32, k int32) int32 {
	i := 0
	for i < len(a) {
		if i > len(a)-2 {
			break
		}
		mul := a[i] * a[i+1]
		if mul > k {
			i++
			continue
		}
		a = append(a[:i], a[i+2:]...)
		a = append(a[:i+1], a[i:]...)
		a[i] = mul
		i = 0
	}

	return int32(len(a))
}

func main() {
	fmt.Println(getMinLength([]int32{1,2,3,4,5,6}, 9))
}
