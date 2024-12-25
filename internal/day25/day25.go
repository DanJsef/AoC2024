package day25

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type lock struct {
	shape [5]int
	size  int
}

type key struct {
	shape [5]int
	size  int
}

func parseInput() (keys []key, locks []lock) {
	file, err := os.Open("./inputs/day25.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for line, err := reader.ReadString('\n'); err == nil; line, err = reader.ReadString('\n') {
		if line == "\n" {
			continue
		}

		temp := [][]rune{}

		for line != "\n" && err == nil {
			temp = append(temp, []rune(line[:len(line)-1]))
			line, err = reader.ReadString('\n')
		}

		shape := [5]int{-1, -1, -1, -1, -1}
		size := -5

		for _, row := range temp {
			for i, r := range row {
				if r == '#' {
					shape[i]++
					size++
				}
			}
		}

		if temp[0][0] == '#' {
			locks = append(locks, lock{shape: shape, size: size})
		} else {
			keys = append(keys, key{shape: shape, size: size})
		}

	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i].size > keys[j].size
	})

	sort.Slice(locks, func(i, j int) bool {
		return locks[i].size > locks[j].size
	})

	return
}

func countCombinations(keys []key, locks []lock) (count int) {
	for _, key := range keys {
		for _, lock := range locks {
			if key.size+lock.size > 25 {
				continue
			}

			fits := true
			for i := 0; i < 5; i++ {
				if key.shape[i]+lock.shape[i] > 5 {
					fits = false
					break
				}
			}

			if fits {
				count++
			}
		}
	}

	return
}

func Run() {
	keys, locks := parseInput()

	fmt.Println("Part one:", countCombinations(keys, locks))
}
