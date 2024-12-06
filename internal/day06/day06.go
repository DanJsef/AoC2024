package day06

import (
	"bufio"
	"fmt"
	"os"
)

type position struct {
	x int
	y int
}

var directons = map[int]position{
	0: {0, -1},
	1: {1, 0},
	2: {0, 1},
	3: {-1, 0},
}

type guard struct {
	x         int
	y         int
	direction int // 0 = north, 1 = east, 2 = south, 3 = west
	finished  bool
}

type floorPlan struct {
	uniqueVisited int
	plan          [][]rune
	finished      int
	guards        []*guard
	initial       position
}

func (fp *floorPlan) simulateStep() {
	for _, guard := range fp.guards {
		guard.x += directons[guard.direction].x
		guard.y += directons[guard.direction].y
		if guard.x < 0 || guard.x >= len(fp.plan[0]) || guard.y < 0 || guard.y >= len(fp.plan) {
			guard.finished = true
			fp.finished++
			continue
		}
		if fp.plan[guard.y][guard.x] == '#' {
			guard.x -= directons[guard.direction].x
			guard.y -= directons[guard.direction].y
			guard.direction = (guard.direction + 1) % 4
			continue
		}
		if fp.plan[guard.y][guard.x] == '.' {
			fp.plan[guard.y][guard.x] = 'x'
			fp.uniqueVisited++
			continue
		}
	}
}

func parseIput() *floorPlan {
	file, err := os.Open("./inputs/day06.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	fp := &floorPlan{}

	i, j := 0, 0

	for input, _, err := reader.ReadRune(); err == nil; input, _, err = reader.ReadRune() {
		if input == '\n' {
			i++
			j = 0
			continue
		}

		if input == '^' {
			fp.uniqueVisited++
			fp.guards = append(fp.guards, &guard{x: j, y: i, direction: 0})
			fp.initial = position{x: j, y: i}
			input = 'x'
		}

		if i == len(fp.plan) {
			fp.plan = append(fp.plan, []rune{})
		}

		fp.plan[i] = append(fp.plan[i], input)

		j++
	}

	return fp
}

func (fp *floorPlan) reset() {
	fp.finished = 0
	fp.guards[0].x = fp.initial.x
	fp.guards[0].y = fp.initial.y
	fp.guards[0].direction = 0
}

func Run() {
	fp := parseIput()

	for fp.finished < len(fp.guards) {
		fp.simulateStep()
	}

	fmt.Println("PartOne: ", fp.uniqueVisited)

	fp.reset()

	loopCount := 0

	for i := 0; i < len(fp.plan); i++ {
		for j := 0; j < len(fp.plan[i]); j++ {

			if i == fp.initial.y && j == fp.initial.x {
				continue
			}
			if fp.plan[i][j] == '#' {
				continue
			}

			fp.plan[i][j] = '#'

			loopCheck := map[string]bool{}
			for fp.finished < len(fp.guards) {
				fp.simulateStep()
				if _, ok := loopCheck[fmt.Sprintf("%d,%d,%d", fp.guards[0].x, fp.guards[0].y, fp.guards[0].direction)]; ok {
					loopCount++
					break
				}

				loopCheck[fmt.Sprintf("%d,%d,%d", fp.guards[0].x, fp.guards[0].y, fp.guards[0].direction)] = true
			}

			fp.plan[i][j] = 'x'
			fp.reset()

		}
	}

	fmt.Println("PartTwo: ", loopCount)
}
