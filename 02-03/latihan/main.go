package main

import "fmt"

func MxDifLg(a1 []string, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}
	a1Max := len(a1[0])
	a1Min := len(a1[0])
	a2Max := len(a2[0])
	a2Min := len(a2[0])

	for i := 0; i < len(a1); i++ {
		if len(a1[i]) > a1Max {
			a1Max = len(a1[i])
		}
		if len(a1[i]) < a1Min {
			a1Min = len(a1[i])
		}
	}

	for i := 0; i < len(a2); i++ {
		if len(a2[i]) > a2Max {
			a2Max = len(a2[i])
		}
		if len(a2[i]) < a2Min {
			a2Min = len(a2[i])
		}
	}

	dif := a1Max - a2Min
	dif2 := a2Max - a1Min

	if dif2 > dif {
		return dif2
	}

	return dif
}

func main() {
	s1 := []string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}
	s2 := []string{"bbbaaayddqbbrrrv"}
	fmt.Println(MxDifLg(s1, s2))
}
