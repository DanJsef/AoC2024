package day1

import (
	"bufio"
	"fmt"
)

func Day1(reader *bufio.Reader) {
	fmt.Println("Day 1")
	for input, err := reader.ReadString('\n'); err == nil; input, err = reader.ReadString('\n') {
		fmt.Print(input)
	}
}
