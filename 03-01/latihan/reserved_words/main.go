package main

import (
	"fmt"
	"strings"
)

func ReverseWords(str string) string {
	strSeq := strings.Split(str, " ")
	result := ""

	for i := len(strSeq) - 1; i >= 0; i-- {
		result = fmt.Sprintf("%s %s", result, strSeq[i])
	}

	return strings.Trim(result, " ")
}

func main() {
	fmt.Println(ReverseWords("hello world!"))
}