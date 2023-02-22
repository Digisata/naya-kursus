package main

import (
	"fmt"
	"strings"
)

func solve(str string) string {
	lowerCount, upperCount := 0, 0
	res := []string{}

	for _, val := range str {
		if val > 90 {
			lowerCount++
			continue
		}

		upperCount++
	}

	if lowerCount == upperCount || lowerCount > upperCount {
		for _, val := range str {
			if val <= 90 {
				val += 32
			}

			res = append(res, string(val))
		}

		return strings.Join(res, "")
	}

	for _, val := range str {
		if val > 90 {
			val -= 32
		}

		res = append(res, string(val))
	}

	return strings.Join(res, "")
}

func main() {
	fmt.Println(solve("a"))
	fmt.Println(solve("Z"))
	fmt.Println(solve("coDe"))
	fmt.Println(solve("CODe"))
	fmt.Println(solve("coDE"))
	fmt.Println(solve("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"))
	strings.ToLower("dfeaw")
}
