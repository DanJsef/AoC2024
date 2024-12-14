package day14

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	datastructs "github.com/DanJsef/AoC2024/internal/data_structs"
	"github.com/DanJsef/AoC2024/internal/utils"
)

type robot struct {
	position datastructs.Position
	velocity datastructs.Position
}

func joinRunes(runes [101]rune) string {
	var builder strings.Builder
	for _, r := range runes {
		if r == 0 {
			builder.WriteRune(' ')
		} else {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

func renderNextStep(robots []*robot) {
	var output [103][101]rune
	for _, robot := range robots {
		robot.simulate(1, 101, 103, nil, nil)
		output[robot.position.Y][robot.position.X] = 'X'
	}
	for i := 0; i < 103; i++ {
		fmt.Println(joinRunes(output[i]))
	}
}

func (r *robot) simulate(seconds int, width int, height int, wg *sync.WaitGroup, ch chan int) {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	for i := 0; i < seconds; i++ {
		r.position = r.position.AddWrap(r.velocity, width, height)
	}

	if ch == nil {
		return
	}

	widthQuadBorder := (width / 2)
	heightQuadBorder := (height / 2)

	if r.position.X == widthQuadBorder || r.position.Y == heightQuadBorder {
		return
	}

	if r.position.X < widthQuadBorder {
		if r.position.Y < heightQuadBorder {
			ch <- 1
			return
		}
		ch <- 3
		return
	}
	if r.position.Y < heightQuadBorder {
		ch <- 2
		return
	}
	ch <- 4
}

func parseInput() []*robot {
	file, err := os.Open("./inputs/day14.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	robots := []*robot{}

	for line, err := reader.ReadString('\n'); err == nil; line, err = reader.ReadString('\n') {
		numbers := utils.ExtractNumbers(line)

		robots = append(robots, &robot{datastructs.Position{X: numbers[0], Y: numbers[1]}, datastructs.Position{X: numbers[2], Y: numbers[3]}})
	}

	return robots
}

func Run() {
	robots := parseInput()

	ch := make(chan int)
	wg := sync.WaitGroup{}

	go func() {
		defer close(ch)
		for _, robot := range robots {
			wg.Add(1)
			robot.simulate(100, 101, 103, &wg, ch)
		}
		wg.Wait()
	}()

	quadrantCounts := make(map[int]int)

	for val := range ch {
		quadrantCounts[val]++
	}

	fmt.Println(quadrantCounts)

	partOne := 1

	for _, v := range quadrantCounts {
		partOne *= v
	}

	fmt.Println("Part one:", partOne)

	robots = parseInput()

	for i := 1; i < 10000; i++ {
		fmt.Println(i)
		renderNextStep(robots)
	}
}
