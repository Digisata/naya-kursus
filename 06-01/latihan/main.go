package main

import "fmt"

type Tuple struct {
	Char  rune
	Count int
}

func OrderedCount(text string) []Tuple {
	res := []Tuple{}
	keys := map[rune]int{}
	for _, val := range text {
		if _, ok := keys[val]; !ok {
			keys[val] = 0
			for _, val2 := range text {
				if val == val2 {
					keys[val]++
				}
			}
			res = append(res, Tuple{Char: val, Count: keys[val]})
		}
	}
	return res
}

func main() {
	fmt.Println(OrderedCount("abracadabra"))
	fmt.Println(OrderedCount("Code Wars"))
	fmt.Println(OrderedCount(""))
}
