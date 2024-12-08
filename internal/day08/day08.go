package day08

import (
	"bufio"
	"fmt"
	"os"
)

type position struct {
	x int
	y int
}

func (p *position) isWithin(pl *plan) bool {
	return p.x >= 0 && p.x < pl.width && p.y >= 0 && p.y < pl.height
}

func (p *position) add(v vector) position {
	return position{x: p.x + v.x, y: p.y + v.y}
}

func (p *position) substract(v vector) position {
	return position{x: p.x - v.x, y: p.y - v.y}
}

type vector struct {
	x int
	y int
}

type antinodes map[[2]int]bool

type antenas map[rune][]position

type plan struct {
	ant         antenas
	anti        antinodes
	antiPartTwo antinodes
	width       int
	height      int
}

func (p *plan) computeAntinodes() {
	for _, antenaVariant := range p.ant {
		for idx, antenaPos := range antenaVariant {
			for _, antenaPos2 := range antenaVariant[idx+1:] {
				vec := vector{x: antenaPos2.x - antenaPos.x, y: antenaPos2.y - antenaPos.y}

				antinodePos := antenaPos.substract(vec)

				antinodePos2 := antenaPos2.add(vec)

				if antinodePos.isWithin(p) {
					p.anti[[2]int{antinodePos.x, antinodePos.y}] = true
				}
				if antinodePos2.isWithin(p) {
					p.anti[[2]int{antinodePos2.x, antinodePos2.y}] = true
				}

				for antinodePos = antenaPos2.substract(vec); antinodePos.isWithin(p); antinodePos = antinodePos.substract(vec) {
					p.antiPartTwo[[2]int{antinodePos.x, antinodePos.y}] = true
				}

				for antinodePos2 = antenaPos.add(vec); antinodePos2.isWithin(p); antinodePos2 = antinodePos2.add(vec) {
					p.antiPartTwo[[2]int{antinodePos2.x, antinodePos2.y}] = true
				}
			}
		}
	}
}

func parseInput() *plan {
	file, err := os.Open("./inputs/day08.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	plan := plan{ant: make(antenas), anti: make(antinodes), antiPartTwo: make(antinodes)}
	i, j := 0, 0

	for input, _, err := reader.ReadRune(); err == nil; input, _, err = reader.ReadRune() {
		if input == '\n' {
			plan.width = j
			i++
			j = 0
			continue
		}

		if input == '.' {
			j++
			continue
		}

		plan.ant[input] = append(plan.ant[input], position{x: j, y: i})
		j++
	}

	plan.height = i

	return &plan
}

func Run() {
	radarPlan := parseInput()

	radarPlan.computeAntinodes()

	fmt.Println("Part one:", len(radarPlan.anti))
	fmt.Println("Part two:", len(radarPlan.antiPartTwo))
}
