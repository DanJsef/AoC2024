package day13

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	datastructs "github.com/DanJsef/AoC2024/internal/data_structs"
)

type machine struct {
	buttonA datastructs.Position
	buttonB datastructs.Position
	prize   datastructs.Position
}

func (m *machine) solve(offset int) int {
	b := (m.buttonA.Y*(m.prize.X+offset) - m.buttonA.X*(m.prize.Y+offset)) / (m.buttonA.Y*m.buttonB.X - m.buttonA.X*m.buttonB.Y)

	a := ((m.prize.X + offset) - m.buttonB.X*b) / m.buttonA.X

	if a*m.buttonA.X+b*m.buttonB.X != (m.prize.X+offset) || a*m.buttonA.Y+b*m.buttonB.Y != (m.prize.Y+offset) {
		return 0
	}

	return a*3 + b
}

func extractNumbers(s string) []int {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(s, -1)

	var numbers []int
	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			fmt.Println("Not a number", match)
		}
		numbers = append(numbers, num)
	}

	return numbers
}

func parseInput() []*machine {
	file, err := os.Open("./inputs/day13.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	i := 0

	machines := []*machine{}

	for line, err := reader.ReadString('\n'); err == nil; line, err = reader.ReadString('\n') {

		if i == 3 {
			i = 0
			continue
		}

		numbers := extractNumbers(line)

		if i == 0 {
			machines = append(machines, &machine{})
			machines[len(machines)-1].buttonA = datastructs.Position{X: numbers[0], Y: numbers[1]}
			i++
			continue
		}

		if i == 1 {
			machines[len(machines)-1].buttonB = datastructs.Position{X: numbers[0], Y: numbers[1]}
			i++
			continue
		}

		if i == 2 {
			machines[len(machines)-1].prize = datastructs.Position{X: numbers[0], Y: numbers[1]}
			i++
			continue
		}

	}

	return machines
}

func Run() {
	machines := parseInput()

	partOne := 0
	partTwo := 0

	for _, m := range machines {
		partOne += m.solve(0)
		partTwo += m.solve(10000000000000)
	}

	fmt.Println("Part one:", partOne)
	fmt.Println("Part two:", partTwo)
}
