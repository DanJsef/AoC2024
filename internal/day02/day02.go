package day02

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
)

type Result struct {
	safe         bool
	safeDampener bool
}

func checkSafe(arr []int) bool {
	increasing := true
	decreasing := true

	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			decreasing = false
		}
		if arr[i] < arr[i-1] {
			increasing = false
		}
		if abs := math.Abs(float64(arr[i] - arr[i-1])); abs > 3 || abs < 1 {
			return false
		}
	}

	return increasing || decreasing
}

func remove(slice []int, s int) []int {
	newSlice := make([]int, len(slice)-1)
	copy(newSlice, slice[:s])
	copy(newSlice[s:], slice[s+1:])
	return newSlice
}

func checkSafeDampener(arr []int) bool {
	if checkSafe(arr) {
		return true
	}

	for i := 0; i < len(arr); i++ {
		if checkSafe(remove(arr, i)) {
			return true
		}
	}

	return false
}

func processLine(input string, ch chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	split := strings.Fields(input)
	intSlice := make([]int, len(split))

	result := Result{true, true}

	for idx, level := range split {
		intSlice[idx], _ = strconv.Atoi(level)
	}

	result.safe = checkSafe(intSlice)
	result.safeDampener = checkSafeDampener(intSlice)

	ch <- result
}

func Run(reader *bufio.Reader) {
	ch := make(chan Result)
	var wg sync.WaitGroup

	go func() {

		for input, err := reader.ReadString('\n'); err == nil; input, err = reader.ReadString('\n') {
			wg.Add(1)
			go processLine(input, ch, &wg)
		}
		defer close(ch)
		wg.Wait()
	}()

	safeCount := 0
	safeCountDampener := 0

	for res := range ch {
		if res.safe {
			safeCount++
		}
		if res.safeDampener {
			safeCountDampener++
		}
	}

	fmt.Println("Safe count:", safeCount)
	fmt.Println("Safe count dampener:", safeCountDampener)
}
