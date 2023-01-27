package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FakeBin(x string) string {
	result := []string{}
	for i, _ := range x {
		num, _ := strconv.Atoi(string(x[i]))
		if num < 5 {
			result = append(result, "0")
			continue
		}

		result = append(result, "1")
	}

	return strings.Join(result, "")
}

func main() {
	fmt.Println(FakeBin("45385593107843568"))
}
