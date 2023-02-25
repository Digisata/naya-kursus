package main

import "fmt"

func GroupArray(arr []int, c int) (res [][]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i != j && arr[i]+arr[j] == c {
				res = append(res, []int{arr[i], arr[j]})
			}
		}
	}
	return
}

func main() {
	fmt.Println(GroupArray([]int{1, 2, 3, 4, 5}, 6))
}