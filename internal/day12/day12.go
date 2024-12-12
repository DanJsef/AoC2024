package day12

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	datastructs "github.com/DanJsef/AoC2024/internal/data_structs"
)

var directions = []datastructs.Position{
	{X: 0, Y: -1},
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: -1, Y: 0},
}

type gardenMap [][]rune

func (gm gardenMap) hasPosition(p datastructs.Position) bool {
	return p.X >= 0 && p.X < len(gm[0]) && p.Y >= 0 && p.Y < len(gm)
}

func (gm gardenMap) calculatePrice() (int, int) {
	visited := make(map[[2]int]bool)

	price := 0
	price2 := 0

	for i := 0; i < len(gm); i++ {
		for j := 0; j < len(gm); j++ {
			if _, ok := visited[[2]int{j, i}]; ok {
				continue
			}
			visited[[2]int{j, i}] = true
			stack := datastructs.Stack[datastructs.Position]{}
			stack.Push(datastructs.Position{X: j, Y: i})

			area := 0
			perimeter := 0
			sides := 0
			visitedBoundry := make(map[[4]int]bool)

			for currPos, ok := stack.Pop(); ok; currPos, ok = stack.Pop() {
				area++
				for _, dir := range directions {
					pos := datastructs.Position{X: currPos.X + dir.X, Y: currPos.Y + dir.Y}
					if !gm.hasPosition(pos) || gm[pos.Y][pos.X] != gm[currPos.Y][currPos.X] {
						if _, ok := visitedBoundry[[4]int{pos.X, pos.Y, dir.X, dir.Y}]; ok {
							perimeter++
							continue
						}
						perimeter++
						sides++
						visitedBoundry[[4]int{pos.X, pos.Y, dir.X, dir.Y}] = true
						if dir == directions[0] || dir == directions[2] {

							for nextPos := pos.Add(directions[1]); gm.hasPosition(datastructs.Position{Y: nextPos.Y - dir.Y, X: nextPos.X - dir.X}) && gm[nextPos.Y-dir.Y][nextPos.X-dir.X] == gm[currPos.Y][currPos.X]; nextPos = nextPos.Add(directions[1]) {
								if !gm.hasPosition(nextPos) || gm[nextPos.Y][nextPos.X] != gm[currPos.Y][currPos.X] {
									visitedBoundry[[4]int{nextPos.X, nextPos.Y, dir.X, dir.Y}] = true
								} else {
									break
								}
							}
							for nextPos := pos.Add(directions[3]); gm.hasPosition(datastructs.Position{Y: nextPos.Y - dir.Y, X: nextPos.X - dir.X}) && gm[nextPos.Y-dir.Y][nextPos.X-dir.X] == gm[currPos.Y][currPos.X]; nextPos = nextPos.Add(directions[3]) {
								if !gm.hasPosition(nextPos) || gm[nextPos.Y][nextPos.X] != gm[currPos.Y][currPos.X] {
									visitedBoundry[[4]int{nextPos.X, nextPos.Y, dir.X, dir.Y}] = true
								} else {
									break
								}
							}
							continue
						}
						if dir == directions[1] || dir == directions[3] {
							for nextPos := pos.Add(directions[0]); gm.hasPosition(datastructs.Position{Y: nextPos.Y - dir.Y, X: nextPos.X - dir.X}) && gm[nextPos.Y-dir.Y][nextPos.X-dir.X] == gm[currPos.Y][currPos.X]; nextPos = nextPos.Add(directions[0]) {
								if !gm.hasPosition(nextPos) || gm[nextPos.Y][nextPos.X] != gm[currPos.Y][currPos.X] {
									visitedBoundry[[4]int{nextPos.X, nextPos.Y, dir.X, dir.Y}] = true
								} else {
									break
								}
							}
							for nextPos := pos.Add(directions[2]); gm.hasPosition(datastructs.Position{Y: nextPos.Y - dir.Y, X: nextPos.X - dir.X}) && gm[nextPos.Y-dir.Y][nextPos.X-dir.X] == gm[currPos.Y][currPos.X]; nextPos = nextPos.Add(directions[2]) {
								if !gm.hasPosition(nextPos) || gm[nextPos.Y][nextPos.X] != gm[currPos.Y][currPos.X] {
									visitedBoundry[[4]int{nextPos.X, nextPos.Y, dir.X, dir.Y}] = true
								} else {
									break
								}
							}
							continue
						}
						continue
					}
					if _, ok := visited[[2]int{pos.X, pos.Y}]; ok {
						continue
					}
					visited[[2]int{pos.X, pos.Y}] = true
					stack.Push(pos)
				}
			}

			price += area * perimeter
			price2 += area * sides

		}
	}

	return price, price2
}

func parseInput() gardenMap {
	file, err := os.Open("./inputs/day12.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	gm := make(gardenMap, 0)

	for line, err := reader.ReadString('\n'); err == nil; line, err = reader.ReadString('\n') {
		gm = append(gm, []rune(strings.TrimSpace(line)))
	}

	return gm
}

func Run() {
	gm := parseInput()
	price, price2 := gm.calculatePrice()
	fmt.Println("Part one:", price)
	fmt.Println("Part two:", price2)
}
