package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	result := fmt.Sprintf("Helooo %s", strings.TrimSuffix(input, "\n"))
	fmt.Println(result)
}
