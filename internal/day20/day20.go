package day20

import (
	"bufio"
	"fmt"
	"os"

	datastructs "github.com/DanJsef/AoC2024/internal/data_structs"
)

var directions = []datastructs.Position{
	{X: 0, Y: -1},
	{X: 1, Y: 0},
	{X: 0, Y: 1},
	{X: -1, Y: 0},
}

type racetrack struct {
	trackMap  [][]rune
	baseTime  int
	timeMap   map[[2]int]int
	start     datastructs.Position
	end       datastructs.Position
	trackPath []datastructs.Position
}

func (rt *racetrack) timeTrack(pos datastructs.Position, time int) {
	rt.timeMap[[2]int{pos.X, pos.Y}] = time
	rt.trackPath = append(rt.trackPath, pos)

	if rt.trackMap[pos.Y][pos.X] == 'E' {
		rt.baseTime = time
		return
	}

	for _, dir := range directions {
		newPos := pos.Add(dir)
		if _, ok := rt.timeMap[[2]int{newPos.X, newPos.Y}]; !ok && rt.trackMap[newPos.Y][newPos.X] != '#' {
			rt.timeTrack(newPos, time+1)
		}
	}
}

func (rt *racetrack) findCheats(cheatTime int, threshold int) int {
	count := 0
	for i := 0; i < len(rt.trackPath); i++ {
		for j := i + 1; j < len(rt.trackPath); j++ {
			if rt.trackPath[i].MahattanDistance(rt.trackPath[j]) > cheatTime {
				continue
			}

			distance := rt.trackPath[i].MahattanDistance(rt.trackPath[j])
			cheatStartTime := rt.timeMap[[2]int{rt.trackPath[i].X, rt.trackPath[i].Y}]
			cheatEndTime := rt.timeMap[[2]int{rt.trackPath[j].X, rt.trackPath[j].Y}]

			cheatTime := cheatStartTime + (rt.baseTime - cheatEndTime)

			if rt.baseTime-cheatTime-distance >= threshold {
				count++
			}
		}
	}
	return count
}

func parseInput() *racetrack {
	file, err := os.Open("./inputs/day20.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	track := &racetrack{timeMap: map[[2]int]int{}}

	for line, err := reader.ReadString('\n'); err == nil; line, err = reader.ReadString('\n') {
		for i, char := range line {
			if char == 'S' {
				track.start = datastructs.Position{X: i, Y: len(track.trackMap)}
			} else if char == 'E' {
				track.end = datastructs.Position{X: i, Y: len(track.trackMap)}
			}
		}
		track.trackMap = append(track.trackMap, []rune(line[:len(line)-1]))
	}

	return track
}

func Run() {
	track := parseInput()

	track.timeTrack(track.start, 0)

	fmt.Println("Part one:", track.findCheats(2, 100))
	fmt.Println("Part two:", track.findCheats(20, 100))
}
