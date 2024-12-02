package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/DanJsef/AoC2024/internal/day01"
	"github.com/DanJsef/AoC2024/internal/day02"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input")
		return
	}

	dayIdx, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("Invalid day input")
		return
	}

	switch int(dayIdx) {
	case 1:
		day01.Run(reader)
	case 2:
		day02.Run(reader)
	default:
		fmt.Println("Invalid day input")
	}
}
