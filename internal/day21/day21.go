package day21

import (
	"fmt"
	"math"

	datastructs "github.com/DanJsef/AoC2024/internal/data_structs"
	"github.com/DanJsef/AoC2024/internal/utils"
)

var directions = map[rune]datastructs.Position{
	'^': {X: 0, Y: -1},
	'>': {X: 1, Y: 0},
	'v': {X: 0, Y: 1},
	'<': {X: -1, Y: 0},
}

type searchPos struct {
	pos datastructs.Position
	seq string
}

type numpad struct {
	verticies       [][]rune
	edges           map[rune][]rune
	vertexPositions map[rune]datastructs.Position
	memo            map[[2]string]int
}

func newNumpad() *numpad {
	verticies := [][]rune{{'7', '8', '9'}, {'4', '5', '6'}, {'1', '2', '3'}, {'#', '0', 'A'}}
	vertexPositions := map[rune]datastructs.Position{
		'7': {X: 0, Y: 0},
		'8': {X: 1, Y: 0},
		'9': {X: 2, Y: 0},
		'4': {X: 0, Y: 1},
		'5': {X: 1, Y: 1},
		'6': {X: 2, Y: 1},
		'1': {X: 0, Y: 2},
		'2': {X: 1, Y: 2},
		'3': {X: 2, Y: 2},
		'#': {X: 0, Y: 3},
		'0': {X: 1, Y: 3},
		'A': {X: 2, Y: 3},
	}
	edges := map[rune][]rune{
		'A': {'<', '^'},
		'0': {'>', '^'},
		'3': {'<', '^', 'v'},
		'2': {'<', '^', 'v', '>'},
		'1': {'^', '>'},
		'6': {'<', '^', 'v'},
		'5': {'<', '^', 'v', '>'},
		'4': {'>', '^', 'v'},
		'9': {'<', 'v'},
		'8': {'<', 'v', '>'},
		'7': {'v', '>'},
	}

	return &numpad{verticies: verticies, edges: edges, vertexPositions: vertexPositions, memo: make(map[[2]string]int)}
}

func newArrowpad() *numpad {
	verticies := [][]rune{{'#', '^', 'A'}, {'<', 'v', '>'}}
	vertexPositions := map[rune]datastructs.Position{
		'#': {X: 0, Y: 0},
		'^': {X: 1, Y: 0},
		'A': {X: 2, Y: 0},
		'<': {X: 0, Y: 1},
		'v': {X: 1, Y: 1},
		'>': {X: 2, Y: 1},
	}
	edges := map[rune][]rune{
		'A': {'<', 'v'},
		'^': {'>', 'v'},
		'>': {'<', '^'},
		'v': {'<', '^', '>'},
		'<': {'>'},
	}

	return &numpad{verticies: verticies, edges: edges, vertexPositions: vertexPositions, memo: make(map[[2]string]int)}
}

func (n *numpad) solveSeq(seq []rune) []string {
	acc := []string{}

	for i := 0; i < len(seq); i++ {
		start := 'A'

		if i > 0 {
			start = seq[i-1]
		}

		result := n.solvePair(n.vertexPositions[start], seq[i])

		if len(acc) == 0 {
			acc = result
			continue
		}

		acc = utils.StringCombinations(acc, result, "A")
	}

	for i := 0; i < len(acc); i++ {
		acc[i] = acc[i] + "A"
	}
	return acc
}

func (n *numpad) solveSeqImproved(seq []rune, cycles int) int {
	if val, ok := n.memo[[2]string{string(seq), string(cycles)}]; ok {
		return val
	}

	length := 0
	start := 'A'

	for _, char := range seq {
		paths := n.solvePair(n.vertexPositions[start], char)

		if cycles == 0 {
			length += len(paths[0] + "A")
			start = char
			continue
		}

		shortest := math.MaxInt64
		for _, path := range paths {
			solved := n.solveSeqImproved([]rune(path+"A"), cycles-1)
			if solved < shortest {
				shortest = solved
			}
		}
		length += shortest
		start = char
	}

	n.memo[[2]string{string(seq), string(cycles)}] = length
	return length
}

func (n *numpad) solvePair(start datastructs.Position, end rune) []string {
	q := datastructs.Queue[searchPos]{}
	q.Enqueue(searchPos{pos: start, seq: ""})

	shortest := math.MaxInt64

	visited := make(map[[2]int]bool)
	visited[[2]int{start.X, start.Y}] = true

	acc := []string{}

	for curr, ok := q.Dequeue(); ok; curr, ok = q.Dequeue() {
		num := n.verticies[curr.pos.Y][curr.pos.X]

		if num == end {
			if len(curr.seq) > shortest {
				break
			}

			shortest = len(curr.seq)
			acc = append(acc, curr.seq)
		}

		for _, dirSign := range n.edges[num] {
			dir := directions[dirSign]

			visited[[2]int{dir.X, dir.Y}] = true
			q.Enqueue(searchPos{pos: curr.pos.Add(dir), seq: curr.seq + string(dirSign)})
		}
	}

	return acc
}

func solveCodes(sequences []string, cycles int) int {
	num := newNumpad()
	sum := 0
	for _, seq := range sequences {
		first := num.solveSeq([]rune(seq))
		arrow := newArrowpad()

		shortest := math.MaxInt64
		for _, seq := range first {
			solved := arrow.solveSeqImproved([]rune(seq), cycles-1)
			if solved < shortest {
				shortest = solved
			}
		}
		sum += shortest * utils.ExtractNumbers(seq)[0]
	}
	return sum
}

func Run() {
	sequences := []string{"029A", "980A", "179A", "456A", "379A"}

	fmt.Println("Part one:", solveCodes(sequences, 2))
	fmt.Println("Part two:", solveCodes(sequences, 25))
}
