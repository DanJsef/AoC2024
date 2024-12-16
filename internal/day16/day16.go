package day16

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"

	datastructs "github.com/DanJsef/AoC2024/internal/data_structs"
)

type vertex struct {
	cost      int
	pos       datastructs.Position
	direction datastructs.Position
	path      [][2]int
}

type MinHeap []vertex

func (heap MinHeap) Len() int           { return len(heap) }
func (heap MinHeap) Less(i, j int) bool { return heap[i].cost < heap[j].cost }
func (heap MinHeap) Swap(i, j int)      { heap[i], heap[j] = heap[j], heap[i] }
func (heap *MinHeap) Push(x any)        { *heap = append(*heap, x.(vertex)) }
func (heap *MinHeap) Pop() any {
	defer func() { *heap = (*heap)[:len(*heap)-1] }()
	return (*heap)[len(*heap)-1]
}

var directions = []datastructs.Position{
	'^': {X: 0, Y: -1},
	'>': {X: 1, Y: 0},
	'v': {X: 0, Y: 1},
	'<': {X: -1, Y: 0},
}

type maze struct {
	mazeMap      [][]rune
	optimalScore int
}

func (m *maze) dijkstra() (int, int) {

	prioQ := &MinHeap{}

	heap.Push(prioQ, vertex{cost: 0, pos: datastructs.Position{X: 1, Y: len(m.mazeMap) - 2}, direction: directions['>'], path: [][2]int{{1, len(m.mazeMap) - 2}}})

	reachCost := make(map[[4]int]int)

	uniqueNodes := make(map[[2]int]bool)

	for prioQ.Len() > 0 {
		currVertex := heap.Pop(prioQ).(vertex)

		if currVertex.cost > m.optimalScore {
			continue
		}

		if m.mazeMap[currVertex.pos.Y][currVertex.pos.X] == 'E' {
			for _, pos := range currVertex.path {
				uniqueNodes[pos] = true
			}

			m.optimalScore = currVertex.cost
		}

		directions := []datastructs.Position{currVertex.direction, currVertex.direction.RotateClockwise(), currVertex.direction.RotateCounterClockwise()}

		for i, dir := range directions {
			nextPos := currVertex.pos.Add(dir)
			if m.mazeMap[nextPos.Y][nextPos.X] == '#' {
				continue
			}

			cost := currVertex.cost + 1
			if i != 0 {
				cost += 1000
			}

			if savedCost, ok := reachCost[[4]int{nextPos.X, nextPos.Y, dir.X, dir.Y}]; !ok || savedCost >= cost {
				reachCost[[4]int{nextPos.X, nextPos.Y, dir.X, dir.Y}] = cost
				newSlice := make([][2]int, len(currVertex.path))
				copy(newSlice, currVertex.path)
				heap.Push(prioQ, vertex{cost: cost, pos: nextPos, direction: dir, path: append(newSlice, [2]int{nextPos.X, nextPos.Y})})
			}
		}
	}

	return m.optimalScore, len(uniqueNodes)
}

func parseInput() *maze {
	file, err := os.Open("./inputs/day16.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	maze := maze{optimalScore: math.MaxInt}
	for line, err := reader.ReadString('\n'); err == nil; line, err = reader.ReadString('\n') {
		maze.mazeMap = append(maze.mazeMap, []rune(line[:len(line)-1]))
	}

	return &maze
}

func Run() {
	maze := parseInput()

	score, seats := maze.dijkstra()
	fmt.Println("Part one:", score)
	fmt.Println("Part two :", seats)
}
