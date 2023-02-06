package main

import (
	"fmt"
	"sort"
	"strings"
)

func TwoToOne(s1 string, s2 string) string {
	s := strings.Split(s1+s2, "")
	sort.Strings(s)

	i := 0
	for i < len(s) {
		if i > len(s)-2 {
			break
		}

		if s[i] == s[i+1] {
			s = append(s[:i], s[i+1:]...)
			i = 0
			continue
		}
		i++
	}

	return strings.Join(s, "")
}

func main() {
	fmt.Println(TwoToOne("aretheyhere", "yestheyarehere"))
}
