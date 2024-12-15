package day15

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	datastructs "github.com/DanJsef/AoC2024/internal/data_structs"
)

var directions = map[rune]datastructs.Position{
	'^': {X: 0, Y: -1},
	'>': {X: 1, Y: 0},
	'v': {X: 0, Y: 1},
	'<': {X: -1, Y: 0},
}

var boxDirections = map[rune]rune{
	'[': '>',
	']': '<',
}

type warehouseMap [][]rune

func (wm warehouseMap) String() string {
	var str string
	for _, row := range wm {
		str += string(row) + "\n"
	}
	return str
}

func (wm warehouseMap) move(robotPos *datastructs.Position, direction datastructs.Position) {
	nextPos := robotPos.Add(direction)
	newNextPos := nextPos
	for wm[newNextPos.Y][newNextPos.X] == 'O' {
		newNextPos = newNextPos.Add(direction)
	}

	if wm[newNextPos.Y][newNextPos.X] == '#' {
		return
	}

	wm[newNextPos.Y][newNextPos.X] = 'O'
	wm[nextPos.Y][nextPos.X] = '@'
	wm[robotPos.Y][robotPos.X] = '.'
	robotPos.X, robotPos.Y = nextPos.X, nextPos.Y
}

func (wm warehouseMap) moveHorizontal(robotPos *datastructs.Position, direction datastructs.Position) {
	nextPositions := []datastructs.Position{robotPos.Add(direction)}

	for wm[nextPositions[len(nextPositions)-1].Y][nextPositions[len(nextPositions)-1].X] == '[' || wm[nextPositions[len(nextPositions)-1].Y][nextPositions[len(nextPositions)-1].X] == ']' {
		nextPositions = append(nextPositions, nextPositions[len(nextPositions)-1].Add(direction))
	}

	if wm[nextPositions[len(nextPositions)-1].Y][nextPositions[len(nextPositions)-1].X] == '#' {
		return
	}

	for i := len(nextPositions) - 1; i >= 0; i-- {
		prevPos := nextPositions[i].Sub(direction)
		wm[nextPositions[i].Y][nextPositions[i].X] = wm[prevPos.Y][prevPos.X]
	}

	wm[robotPos.Y][robotPos.X] = '.'
	robotPos.X, robotPos.Y = nextPositions[0].X, nextPositions[0].Y
}

func (wm warehouseMap) moveVertical(robotPos *datastructs.Position, direction datastructs.Position) {
	nextPos := robotPos.Add(direction)
	if wm[nextPos.Y][nextPos.X] == '#' {
		return
	}

	canMove := true

	if wm[nextPos.Y][nextPos.X] == '[' || wm[nextPos.Y][nextPos.X] == ']' {
		nextBoxPositions := []datastructs.Position{}
		addedBoxPositions := make(map[[2]int]bool)
		if _, ok := addedBoxPositions[[2]int{nextPos.X, nextPos.Y}]; !ok {
			addedBoxPositions[[2]int{nextPos.X, nextPos.Y}] = true
			nextBoxPositions = append(nextBoxPositions, nextPos)
		}

		partBoxPos := nextPos.Add(directions[boxDirections[wm[nextPos.Y][nextPos.X]]])
		if _, ok := addedBoxPositions[[2]int{partBoxPos.X, partBoxPos.Y}]; !ok {
			addedBoxPositions[[2]int{partBoxPos.X, partBoxPos.Y}] = true
			nextBoxPositions = append(nextBoxPositions, partBoxPos)
		}

		canMove = wm.moveVerticalLevel(nextBoxPositions, direction)
	}

	if canMove {
		wm[nextPos.Y][nextPos.X] = '@'
		wm[robotPos.Y][robotPos.X] = '.'
		robotPos.X, robotPos.Y = nextPos.X, nextPos.Y
	}

}

func (wm warehouseMap) moveVerticalLevel(boxPositions []datastructs.Position, direction datastructs.Position) bool {
	if len(boxPositions) == 0 {
		return true
	}

	nextBoxPositions := []datastructs.Position{}
	addedBoxPositions := make(map[[2]int]bool)

	for _, boxPos := range boxPositions {
		nextBoxPos := boxPos.Add(direction)
		if wm[nextBoxPos.Y][nextBoxPos.X] == '#' {
			return false
		}

		if wm[nextBoxPos.Y][nextBoxPos.X] == '[' || wm[nextBoxPos.Y][nextBoxPos.X] == ']' {
			if _, ok := addedBoxPositions[[2]int{nextBoxPos.X, nextBoxPos.Y}]; !ok {
				addedBoxPositions[[2]int{nextBoxPos.X, nextBoxPos.Y}] = true
				nextBoxPositions = append(nextBoxPositions, nextBoxPos)
			}

			partBoxPos := nextBoxPos.Add(directions[boxDirections[wm[nextBoxPos.Y][nextBoxPos.X]]])
			if _, ok := addedBoxPositions[[2]int{partBoxPos.X, partBoxPos.Y}]; !ok {
				addedBoxPositions[[2]int{partBoxPos.X, partBoxPos.Y}] = true
				nextBoxPositions = append(nextBoxPositions, partBoxPos)
			}
		}
	}

	canMove := wm.moveVerticalLevel(nextBoxPositions, direction)

	if !canMove {
		return false
	}

	for _, boxPos := range boxPositions {
		nextBoxPos := boxPos.Add(direction)
		wm[nextBoxPos.Y][nextBoxPos.X] = wm[boxPos.Y][boxPos.X]
		wm[boxPos.Y][boxPos.X] = '.'
	}

	return true
}

func (wm warehouseMap) sumBoxGps(edge rune) int {
	sum := 0

	for y, row := range wm {
		for x, r := range row {
			if r == edge {
				sum += x + (y * 100)
			}
		}
	}

	return sum
}

func parseWarehouse(reader *bufio.Reader) (warehouseMap, datastructs.Position, warehouseMap, datastructs.Position) {
	warehouse := warehouseMap{}
	warehouse2 := warehouseMap{}
	robotPosition := datastructs.Position{}
	robotPosition2 := datastructs.Position{}

	i := 0

	for line, err := reader.ReadString('\n'); line != "\n"; line, err = reader.ReadString('\n') {
		if err != nil {
			fmt.Println("Error reading file:", err)
			return nil, datastructs.Position{}, nil, datastructs.Position{}
		}

		for j, r := range line {
			if r == '@' {
				robotPosition = datastructs.Position{X: j, Y: i}
			}
		}

		warehouse = append(warehouse, []rune(line[:len(line)-1]))

		line = strings.ReplaceAll(line, "#", "##")
		line = strings.ReplaceAll(line, "O", "[]")
		line = strings.ReplaceAll(line, ".", "..")
		line = strings.ReplaceAll(line, "@", "@.")

		for j, r := range line {
			if r == '@' {
				robotPosition2 = datastructs.Position{X: j, Y: i}
			}
		}

		warehouse2 = append(warehouse2, []rune(line[:len(line)-1]))

		i++
	}

	return warehouse, robotPosition, warehouse2, robotPosition2
}

func Run() {
	file, err := os.Open("./inputs/day15.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	warehouse, robotPos, warehouse2, robotPos2 := parseWarehouse(reader)

	for input, _, err := reader.ReadRune(); err == nil; input, _, err = reader.ReadRune() {
		if input == '\n' {
			continue
		}

		warehouse.move(&robotPos, directions[input])

		if input == 'v' || input == '^' {
			warehouse2.moveVertical(&robotPos2, directions[input])
		}
		if input == '>' || input == '<' {
			warehouse2.moveHorizontal(&robotPos2, directions[input])
		}
	}

	fmt.Println("Part one:", warehouse.sumBoxGps('O'))
	fmt.Println("Part two:", warehouse2.sumBoxGps('['))
}
