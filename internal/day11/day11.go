package day11

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func splitNumber(num int) ([]int, bool) {
	digits := int(math.Log10(float64(num))) + 1

	if digits%2 != 0 {
		return []int{0, 0}, false
	}

	halfDigits := digits / 2
	divisor := int(math.Pow(10, float64(halfDigits)))

	part1 := num / divisor
	part2 := num % divisor

	return []int{part1, part2}, true
}

func changeStone(stone int) []int {
	if stone == 0 {
		return []int{1}
	}

	if split, ok := splitNumber(stone); ok {
		return split
	}

	return []int{stone * 2024}
}

func simulator() func(stone int, blinks int) int {
	solved_cache := make(map[[2]int]int)
	change_cache := make(map[int][]int)

	var simulate func(int, int) int
	simulate = func(stone int, blinks int) int {
		if blinks == 0 {
			return 1
		}

		if solved, ok := solved_cache[[2]int{stone, blinks}]; ok {
			return solved
		}

		var change []int
		if change_cached, ok := change_cache[stone]; ok {
			change = change_cached
		} else {
			change = changeStone(stone)
			change_cache[stone] = change
		}

		acc := 0

		acc += simulate(change[0], blinks-1)

		if len(change) == 2 {
			acc += simulate(change[1], blinks-1)
		}

		solved_cache[[2]int{stone, blinks}] = acc

		return acc
	}

	return simulate
}

func simulate(stones []int, blinks int) int {
	simulate := simulator()
	acc := 0
	for _, stone := range stones {
		acc += simulate(stone, blinks)
	}

	return acc
}

func parseInput() []int {
	file, err := os.Open("./inputs/day11.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	input := []int{}

	for line, err := reader.ReadString('\n'); err == nil; line, err = reader.ReadString('\n') {
		split := strings.Fields(line)

		for _, s := range split {
			num, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println("Not a number:", s)
			}
			input = append(input, num)
		}
	}

	return input
}

func Run() {
	input := parseInput()

	fmt.Println("Part one:", simulate(input, 25))
	fmt.Println("Part two:", simulate(input, 75))
}
