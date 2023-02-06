package main

import (
	"fmt"
	"strings"
)

func FindShort(s string) int {
	wordSeq := strings.Split(s, " ")
	shortest := len(wordSeq[0])

	for _, word := range wordSeq {
		if len(word) < shortest {
			shortest = len(word)
		}
	}

	return shortest
}

func main() {
	fmt.Println(FindShort("bitcoin take over the world maybe who knows perhaps"))
}
