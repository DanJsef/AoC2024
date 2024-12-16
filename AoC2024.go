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
	"github.com/DanJsef/AoC2024/internal/day09"
	"github.com/DanJsef/AoC2024/internal/day10"
	"github.com/DanJsef/AoC2024/internal/day11"
	"github.com/DanJsef/AoC2024/internal/day12"
	"github.com/DanJsef/AoC2024/internal/day13"
	"github.com/DanJsef/AoC2024/internal/day14"
	"github.com/DanJsef/AoC2024/internal/day15"
	"github.com/DanJsef/AoC2024/internal/day16"
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
	case 9:
		day09.Run()
	case 10:
		day10.Run()
	case 11:
		day11.Run()
	case 12:
		day12.Run()
	case 13:
		day13.Run()
	case 14:
		day14.Run()
	case 15:
		day15.Run()
	case 16:
		day16.Run()
	default:
		fmt.Println("Invalid day input")
	}
}
