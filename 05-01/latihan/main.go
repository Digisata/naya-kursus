package main

import (
	"fmt"
	"strings"
)

func solution(str, ending string) bool {
	str = strings.Trim(str, " ")
	ending = strings.Trim(ending, " ")
	if len(str) < len(ending) {
		return false
	}
	return str[len(str) - len(ending):] == ending
}

func main() {
	fmt.Println(solution("", ""))
	fmt.Println(solution("a", "z"))
}