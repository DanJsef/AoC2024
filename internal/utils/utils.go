package utils

import (
	"fmt"
	"math"
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

func StringCombinations(arr1, arr2 []string, delimeter string) []string {
	acc := []string{}

	for _, s1 := range arr1 {
		for _, s2 := range arr2 {
			acc = append(acc, s1+delimeter+s2)
		}
	}

	return acc
}

func KeepShortestStrings(strings []string) []string {
	if len(strings) == 0 {
		return strings
	}

	shortestLength := math.MaxInt32
	for _, str := range strings {
		if len(str) < shortestLength {
			shortestLength = len(str)
		}
	}

	result := []string{}
	for _, str := range strings {
		if len(str) == shortestLength {
			result = append(result, str)
		}
	}

	return result
}
