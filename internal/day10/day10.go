package day10

import (
	"bufio"
	"fmt"
	"os"

	"github.com/DanJsef/AoC2024/internal/data_structs"
)

type position struct {
	x int
	y int
}

var directons = []position{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

type trailhead struct {
	x      int
	y      int
	score  int
	rating int
}

type trailMap [][]int

func (tm trailMap) isWithin(p position) bool {
	return p.x >= 0 && p.x < len(tm[0]) && p.y >= 0 && p.y < len(tm)
}

func (tm trailMap) traverse(ths []*trailhead) (int, int) {
	finalScore := 0
	finalRating := 0
	for _, th := range ths {
		stack := datastructs.Stack[position]{}

		visited := make(map[[2]int]bool)

		stack.Push(position{th.x, th.y})

		for currPos, ok := stack.Pop(); ok; currPos, ok = stack.Pop() {

			currHeight := tm[currPos.y][currPos.x]

			if currHeight == 9 {
				if !visited[[2]int{currPos.x, currPos.y}] {
					th.score++
				}
				th.rating++
				visited[[2]int{currPos.x, currPos.y}] = true
				continue
			}

			for _, dir := range directons {
				if !tm.isWithin(position{currPos.x + dir.x, currPos.y + dir.y}) {
					continue
				}
				if tm[currPos.y+dir.y][currPos.x+dir.x] != currHeight+1 {
					continue
				}

				stack.Push(position{currPos.x + dir.x, currPos.y + dir.y})
			}

			visited[[2]int{currPos.x, currPos.y}] = true
		}

		finalScore += th.score
		finalRating += th.rating
	}

	return finalScore, finalRating
}

func parseInput() (trailMap, []*trailhead) {
	file, err := os.Open("./inputs/day10.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	tm := make(trailMap, 0)

	ths := make([]*trailhead, 0)

	i, j := 0, 0

	for input, _, err := reader.ReadRune(); err == nil; input, _, err = reader.ReadRune() {
		if input == '\n' {
			i++
			j = 0
			continue
		}
		var height int
		if input >= '0' && input <= '9' {
			height = int(input - '0')
		} else {
			fmt.Println("Not a digit")
		}

		if height == 0 {
			ths = append(ths, &trailhead{x: j, y: i, score: 0})
		}

		if i == len(tm) {
			tm = append(tm, make([]int, 0))
		}

		tm[i] = append(tm[i], height)
		j++
	}

	return tm, ths
}

func Run() {
	tm, ths := parseInput()

  score, rating := tm.traverse(ths)

	fmt.Println("Part one:", score)
	fmt.Println("Part two:", rating)
}
