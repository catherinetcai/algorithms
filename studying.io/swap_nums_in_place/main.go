package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readInput()
	x, _ := strconv.Atoi(strings.TrimSpace(input[0]))
	y, _ := strconv.Atoi(strings.TrimSpace(input[1]))
	fmt.Printf("Before swap: x(%d), y(%d)\n", x, y)
	x = x ^ y
	y = x ^ y
	x = x ^ y
	fmt.Printf("After swap: x(%d), y(%d)\n", x, y)
}

func readInput() []string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter input numbers delimited by a space.")
	input, _ := reader.ReadString('\n')
	return strings.Split(input, " ")
}
