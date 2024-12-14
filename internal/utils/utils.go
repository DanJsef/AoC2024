package utils

import (
	"fmt"
	"regexp"
	"strconv"
)

func ExtractNumbers(s string) []int {
	re := regexp.MustCompile(`-?\d+`)
	matches := re.FindAllString(s, -1)

	var numbers []int
	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			fmt.Println("Not a number", match)
		}
		numbers = append(numbers, num)
	}

	return numbers
}
