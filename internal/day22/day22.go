package day22

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func pasreInput() []int {
	file, err := os.Open("./inputs/day22.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	numbers := []int{}

	for line, err := reader.ReadString('\n'); err == nil; line, err = reader.ReadString('\n') {
		val, err := strconv.Atoi(line[:len(line)-1])
		if err != nil {
			fmt.Println("Error parsing line:", err)
			return nil
		}

		numbers = append(numbers, val)
	}

	return numbers
}

func mix(secret, value int) int {
	return secret ^ value
}

func prune(secret int) int {
	return secret % 16777216
}

func generator() func(secret int) int {
	cache := make(map[int]int)

	generate := func(secret int) int {
		if cached, ok := cache[secret]; ok {
			return cached
		}

		firstStep := secret * 64
		firstStep = mix(secret, firstStep)
		firstStep = prune(firstStep)

		secondStep := firstStep / 32
		secondStep = mix(firstStep, secondStep)
		secondStep = prune(secondStep)

		thirdStep := secondStep * 2048
		thirdStep = mix(secondStep, thirdStep)
		thirdStep = prune(thirdStep)

		cache[secret] = thirdStep
		return thirdStep
	}
	return generate
}

func shiftAdd(arr [4]int, num int) [4]int {
	for i := 0; i < 3; i++ {
		arr[i] = arr[i+1]
	}
	arr[3] = num
	return arr
}

func Run() {

	secrets := pasreInput()

	ranking := make([]map[[4]int]int, len(secrets))
	uniqueSeq := make(map[[4]int]bool)

	generate := generator()

	sum := 0

	for i, secret := range secrets {
		lastSeq := [4]int{}
		previousPrice := 0
		ranking[i] = make(map[[4]int]int)
		for j := 0; j < 2000; j++ {
			secret = generate(secret)

			price := secret % 10

			if j < 4 {
				lastSeq[j] = price - previousPrice
			} else {
				lastSeq = shiftAdd(lastSeq, price-previousPrice)
				if _, ok := ranking[i][lastSeq]; !ok {
					ranking[i][lastSeq] = price
					uniqueSeq[lastSeq] = true
				}
			}
			previousPrice = price
		}

		sum += secret
	}

	fmt.Println("Part one:", sum)

	bestSum := 0

	for seq := range uniqueSeq {
		sum := 0
		for rank := range ranking {
			if price, ok := ranking[rank][seq]; ok {
				sum += price
			}
		}
		if sum > bestSum {
			bestSum = sum
		}
	}

	fmt.Println("Part two:", bestSum)

}
