package day19

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type node struct {
	end  bool
	next map[rune]*node
}

type trie map[rune]*node

var cache = make(map[string]int)

func (t trie) insert(word string) {
	currNode := t
	for i, char := range word {
		if _, ok := currNode[char]; !ok {
			currNode[char] = &node{next: make(map[rune]*node)}
		}
		if i == len(word)-1 {
			currNode[char].end = true
		}
		currNode = currNode[char].next
	}
}

func (t trie) findSolution(word string) bool {
	toTryIndexes := []int{}
	curr := t
	for i, char := range word {
		match, ok := curr[char]
		if !ok {

			break
		}
		if match.end {
			toTryIndexes = append(toTryIndexes, i)
			if i == len(word)-1 {
				return true
			}
		}
		curr = match.next
	}

	for _, index := range toTryIndexes {
		if t.findSolution(word[index+1:]) {
			return true
		}
	}

	return false
}

func (t trie) findAllSolutions(word string) int {
	if val, ok := cache[word]; ok {
		return val
	}

	toTryIndexes := []int{}
	curr := t

	found := 0
	for i, char := range word {
		match, ok := curr[char]
		if !ok {

			break
		}
		if match.end {
			toTryIndexes = append(toTryIndexes, i)
			if i == len(word)-1 {
				found++
			}
		}
		curr = match.next
	}

	for _, index := range toTryIndexes {
		found += t.findAllSolutions(word[index+1:])
	}

	cache[word] = found

	return found
}

func parseInput() (trie, []string) {
	file, err := os.Open("./inputs/day19.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	t := trie{}
	designs := []string{}

	i := 0

	for line, err := reader.ReadString('\n'); err == nil; line, err = reader.ReadString('\n') {
		line = line[:len(line)-1]
		if i == 0 {
			for _, towel := range strings.Split(line, ", ") {
				t.insert(towel)
			}
			i++
			continue
		}
		if i == 1 {
			i++
			continue
		}

		designs = append(designs, line)
	}

	return t, designs
}

func Run() {
	t, designs := parseInput()

	partOne := 0
	partTwo := 0

	for _, d := range designs {
		if t.findSolution(d) {
			partOne++
		}
		partTwo += t.findAllSolutions(d)
	}

	fmt.Println("Part one:", partOne)
	fmt.Println("Part two:", partTwo)
}
