package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for {
		fmt.Print("Enter a number: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		number, err := strconv.Atoi(strings.TrimSuffix(input, "\n"))
		if err != nil {
			fmt.Println("Please enter a number!")
			continue
		}
		fmt.Println(fact(number))
		return
	}
}

func fact(n int) int {
	if n == 1 {
		return 1
	}

	return n * fact(n-1)
}
