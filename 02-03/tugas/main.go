package main

import "fmt"

func romanToInt(s string) (res int) {
	roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	i := 0
	for i < len(s) {
		if i != len(s)-1 && roman[string(s[i])] < roman[string(s[i+1])] {
			res += roman[string(s[i+1])] - roman[string(s[i])]
			if i+2 > len(s)-1 {
				return
			}
			i += 2
			continue
		} else {
			res += roman[string(s[i])]
			i++
		}
	}
	return
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
