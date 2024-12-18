package day18

import (
	"bufio"
	"fmt"
	"os"

	datastructs "github.com/DanJsef/AoC2024/internal/data_structs"
	"github.com/DanJsef/AoC2024/internal/utils"
)

type memoryMap [71][71]rune
type corruptList []datastructs.Position

type searchPos struct {
	pos    datastructs.Position
	length int
}

var directions = []datastructs.Position{
	{X: 0, Y: -1},
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: -1, Y: 0},
}

func (mm memoryMap) bfs() int {
	q := datastructs.Queue[searchPos]{}
	q.Enqueue(searchPos{pos: datastructs.Position{X: 0, Y: 0}, length: 0})

	visited := make(map[[2]int]bool)

	for curr, ok := q.Dequeue(); ok; curr, ok = q.Dequeue() {
		if curr.pos.X == len(mm)-1 && curr.pos.Y == len(mm)-1 {
			return curr.length
		}

		for _, dir := range directions {
			nextPos := curr.pos.Add(dir)
			if !nextPos.IsWithinBounds(len(mm), len(mm)) || visited[[2]int{nextPos.X, nextPos.Y}] || mm[nextPos.Y][nextPos.X] == '#' {
				continue
			}

			visited[[2]int{nextPos.X, nextPos.Y}] = true
			q.Enqueue(searchPos{pos: nextPos, length: curr.length + 1})
		}
	}

	return 0
}

func parseInput() corruptList {
	file, err := os.Open("./inputs/day18.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	corrupt := corruptList{}
	for line, err := reader.ReadString('\n'); err == nil; line, err = reader.ReadString('\n') {
		numbers := utils.ExtractNumbers(line)

		corrupt = append(corrupt, datastructs.Position{X: numbers[0], Y: numbers[1]})

	}
	return corrupt
}

func Run() {
	corrupt := parseInput()

	memory := memoryMap{}
	i := 0
	for ; i < 1023; i++ {
		memory[corrupt[i].Y][corrupt[i].X] = '#'
	}
	fmt.Println("Part one:", memory.bfs())

	//Improve to binary search if input was longer?
	for ; i < len(corrupt); i++ {
		memory[corrupt[i].Y][corrupt[i].X] = '#'

		if memory.bfs() == 0 {
			fmt.Printf("Part two: %d,%d\n", corrupt[i].X, corrupt[i].Y)
			return
		}
	}
}
