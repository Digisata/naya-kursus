package main

import "fmt"

func NbYear(p0 int, percent float64, aug int, p int) (res int) {
	pop := p0
	for {
		pop = pop + int(float64(pop)*percent/100) + aug
		res++
		if pop >= p {
			return
		}
	}
}

func main() {
	fmt.Println(NbYear(1500, 5, 100, 5000))
	fmt.Println(NbYear(1500000, 2.5, 10000, 2000000))
}
