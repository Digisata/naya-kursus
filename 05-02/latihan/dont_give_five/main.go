package main

import "fmt"

func DontGiveMeFive(start int, end int) (counter int) {
OuterLoop:
	for i := start; i <= end; i++ {
		num := i
		if num < 0 {
			num = -num
		}
		for num > 0 {
			if num%10 == 5 {
				continue OuterLoop
			}

			num = num / 10
		}
		counter++
	}
	return
}

func main() {
	fmt.Println(DontGiveMeFive(1, 9))
}
