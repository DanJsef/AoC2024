package day01

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type MinHeap []int

func (heap MinHeap) Len() int           { return len(heap) }
func (heap MinHeap) Less(i, j int) bool { return heap[i] < heap[j] }
func (heap MinHeap) Swap(i, j int)      { heap[i], heap[j] = heap[j], heap[i] }
func (heap *MinHeap) Push(x any)        { *heap = append(*heap, x.(int)) }
func (heap *MinHeap) Pop() any {
	defer func() { *heap = (*heap)[:len(*heap)-1] }()
	return (*heap)[len(*heap)-1]
}

func Run(reader *bufio.Reader) {
	leftList := &MinHeap{}
	rightList := &MinHeap{}

	rightListCountMap := make(map[int]int)

	for input, err := reader.ReadString('\n'); err == nil; input, err = reader.ReadString('\n') {

		split := strings.Fields(input)

		leftValue, _ := strconv.Atoi(split[0])
		rightValue, _ := strconv.Atoi(split[1])
		heap.Push(leftList, leftValue)
		heap.Push(rightList, rightValue)

		rightListCountMap[rightValue]++
	}

	totalDistance := 0

	similarityScore := 0

	for leftList.Len() > 0 {
		leftPop := heap.Pop(leftList).(int)
		rightPop := heap.Pop(rightList).(int)
		totalDistance += int(math.Abs(float64(leftPop - rightPop)))

		similarityScore += leftPop * rightListCountMap[leftPop]
	}

	fmt.Println("Total distance:", totalDistance)
	fmt.Println("Similarity score:", similarityScore)
}
