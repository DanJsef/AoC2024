package day07

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

func parseInput() [][]int {
	file, err := os.Open("./inputs/day07.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	input := [][]int{}

	for line, err := reader.ReadString('\n'); err == nil; line, err = reader.ReadString('\n') {
		line = strings.TrimSpace(line)
		split := strings.Split(line, ":")

		equationParts := strings.Fields(split[1])

		equationSlice := make([]int, len(equationParts)+1)

		equationResut, err := strconv.Atoi(split[0])
		if err != nil {
			fmt.Println("Not a number:", equationParts[0])
		}

		equationSlice[0] = equationResut

		for i, part := range equationParts {
			intPart, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Not a number:", intPart)
			}
			equationSlice[i+1] = intPart
		}

		input = append(input, equationSlice)
	}

	return input
}

func concatenateIntegers(x, y int) int {
	length := int(math.Log10(float64(y))) + 1

	return x*int(math.Pow(10, float64(length))) + y
}

func solver(equation []int, concatationEnabled bool) func(idx int, acc int) bool {
	var solve func(idx int, acc int) bool
	solve = func(idx int, acc int) bool {
		if idx == len(equation) {
			return acc == equation[0]
		}

		if acc > equation[0] {
			return false
		}

		if concatationEnabled {
			concat := solve(idx+1, concatenateIntegers(acc, equation[idx]))
			if concat {
				return concat
			}
		}

		multiply := solve(idx+1, acc*equation[idx])
		if multiply {
			return multiply
		}

		add := solve(idx+1, acc+equation[idx])
		if add {
			return add
		}

		return false
	}
	return solve
}

func solvePart(equation []int, ch chan int, wg *sync.WaitGroup, partTwo bool) {
	defer wg.Done()
	solve := solver(equation, partTwo)
	result := solve(2, equation[1])
	if result {
		ch <- equation[0]
	}
}

func Run() {
	input := parseInput()

	partOneSum := 0

	results := make(chan int)
	var wg sync.WaitGroup

	go func() {
		for _, equation := range input {
			wg.Add(1)
			go solvePart(equation, results, &wg, false)
		}
		defer close(results)
		wg.Wait()
	}()

	for res := range results {
		partOneSum += res
	}

	fmt.Println("Part one:", partOneSum)

	partTwoSum := 0

	results = make(chan int)

	go func() {
		for _, equation := range input {
			wg.Add(1)
			go solvePart(equation, results, &wg, true)
		}
		defer close(results)
		wg.Wait()
	}()

	for res := range results {
		partTwoSum += res
	}

	fmt.Println("Part two:", partTwoSum)
}
