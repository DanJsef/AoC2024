package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/DanJsef/AoC2024/internal/day01"
	"github.com/DanJsef/AoC2024/internal/day02"
	"github.com/DanJsef/AoC2024/internal/day03"
	"github.com/DanJsef/AoC2024/internal/day04"
	"github.com/DanJsef/AoC2024/internal/day05"
	"github.com/DanJsef/AoC2024/internal/day06"
	"github.com/DanJsef/AoC2024/internal/day07"
	"github.com/DanJsef/AoC2024/internal/day08"
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
	case 3:
		day03.Run()
	case 4:
		day04.Run()
	case 5:
		day05.Run()
	case 6:
		day06.Run()
	case 7:
		day07.Run()
	case 8:
		day08.Run()
	default:
		fmt.Println("Invalid day input")
	}
}
