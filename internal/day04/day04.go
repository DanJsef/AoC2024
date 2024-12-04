package day04

import (
	"bufio"
	"fmt"
	"os"
)

type position struct {
	x int
	y int
}

type wordSearch [][]rune

const matchWord = "XMAS"

var directons = map[string]position{
	"north":     {0, 1},
	"northwest": {1, 1},
	"west":      {1, 0},
	"southwest": {1, -1},
	"south":     {0, -1},
	"southeast": {-1, -1},
	"east":      {-1, 0},
	"northeast": {-1, 1},
}

func prepareWordSearch() wordSearch {
	file, _ := os.Open("./inputs/day04.txt")
	defer file.Close()

	reader := bufio.NewReader(file)

	var ws wordSearch
	for line, err := reader.ReadString('\n'); err == nil; line, err = reader.ReadString('\n') {
		ws = append(ws, []rune(line))
	}
	return ws
}

func scanDirection(ws *wordSearch, pos position, direction position) bool {

	for lenght := 0; lenght < len(matchWord); lenght++ {
		scanPos := position{pos.x + (lenght * direction.x), pos.y + (lenght * direction.y)}

		if scanPos.x < 0 || scanPos.x >= len(*ws) || scanPos.y < 0 || scanPos.y >= len((*ws)[scanPos.x]) {
			return false
		}

		if (*ws)[scanPos.x][scanPos.y] != rune(matchWord[lenght]) {
			return false
		}
	}

	return true
}

func scanPosition(ws *wordSearch, pos position) int {
	found := 0

	for _, direction := range directons {
		if scanDirection(ws, pos, direction) {
			found++
		}
	}

	return found
}

func scanCrossPart(ws *wordSearch, pos position, offsets [2]position) bool {
	sRuneFound := false
	mRuneFound := false

	for _, offset := range offsets {
		sRuneFound = sRuneFound || (*ws)[pos.x+offset.x][pos.y+offset.y] == 'S'
		mRuneFound = mRuneFound || (*ws)[pos.x+offset.x][pos.y+offset.y] == 'M'
	}

	return sRuneFound && mRuneFound
}

func scanPositionX(ws *wordSearch, pos position) bool {
	return (*ws)[pos.x][pos.y] == 'A' && scanCrossPart(ws, pos, [2]position{directons["northeast"], directons["southwest"]}) && scanCrossPart(ws, pos, [2]position{directons["southeast"], directons["northwest"]})

}

func Run() {
	ws := prepareWordSearch()

	fmt.Println(ws)

	totalFoundPartOne := 0
	totalFoundPartTwo := 0

	for x := 0; x < len(ws); x++ {
		for y := 0; y < len(ws[x]); y++ {
			totalFoundPartOne += scanPosition(&ws, position{x, y})

			if x > 0 && x < len(ws)-1 && y > 0 && y < len(ws[x])-1 && scanPositionX(&ws, position{x, y}) {
				totalFoundPartTwo++
			}
		}
	}

	fmt.Println("Total XMAS found:", totalFoundPartOne)
	fmt.Println("Total X-MAS found:", totalFoundPartTwo)
}
