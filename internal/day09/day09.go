package day09

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type linkNode struct {
	fileId int
	idx    int
	size   int
	prev   *linkNode
	next   *linkNode
}

func (l *linkNode) append(n *linkNode) {
	n.idx = l.idx
	l.next.prev = n
	n.next = l.next
	l.next = n
	n.prev = l
}

type diskMap struct {
	first *linkNode
	last  *linkNode
}

func (d *diskMap) reformat() {
	for currNode := d.first; currNode != nil; currNode = currNode.next {
		if currNode.fileId != -1 {
			continue
		}

		moveNode := d.last
		for moveNode.fileId == -1 {
			moveNode.prev.next = nil
			d.last = moveNode.prev
			moveNode = d.last
		}

		if currNode.idx > moveNode.idx {
			break
		}

		currNode.fileId = moveNode.fileId
		moveNode.prev.next = nil
		d.last = moveNode.prev
	}
}

func (d *diskMap) reformatPartTwo() {
	for moveNode := d.last; moveNode != nil; moveNode = moveNode.prev {
		if moveNode.fileId == -1 {
			continue
		}

		currNode := d.first
		for currNode.idx < moveNode.idx {
			if currNode.fileId != -1 {
				currNode = currNode.next
				continue
			}

			if currNode.size < moveNode.size {
				currNode = currNode.next
				continue
			}

			if currNode.size > moveNode.size {
				currNode.append(&linkNode{fileId: -1, size: currNode.size - moveNode.size})
			}

			currNode.fileId = moveNode.fileId
			currNode.size = moveNode.size
			moveNode.fileId = -1
			break
		}
	}
}

func (d *diskMap) checkSum() int {
	i, sum := 0, 0

	for currNode := d.first; currNode != nil; currNode = currNode.next {
		sum += i * currNode.fileId
		i++
	}

	return sum
}

func (d *diskMap) checkSumPartTwo() int {
	i, sum := 0, 0

	for currNode := d.first; currNode != nil; currNode = currNode.next {
		for size := 0; size < currNode.size; size++ {
			localSum := i * currNode.fileId
			if localSum > 0 {
				sum += i * currNode.fileId
			}
			i++
		}
	}

	return sum
}

func (d *diskMap) getVisual(mode int) string {
	res := ""
	for currNode := d.first; currNode != nil; currNode = currNode.next {
		if currNode.fileId == -1 {
			if mode == 1 {
				for i := 0; i < currNode.size; i++ {
					res += fmt.Sprintf(".")
				}
			} else {
				res += fmt.Sprintf(".")
			}
			continue
		}
		if mode == 1 {
			for i := 0; i < currNode.size; i++ {
				res += fmt.Sprintf("%d", currNode.fileId)
			}
		} else {
			res += fmt.Sprintf("%d", currNode.fileId)
		}
	}
	return res
}

func parseInput() (*diskMap, *diskMap) {
	file, err := os.Open("./inputs/day09.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	isFileId := true

	disk := diskMap{}
	disk2 := diskMap{}

	var currentNode *linkNode
	var currentNode2 *linkNode

	idSeq := 0

	idx := 0
	idx2 := 0

	for input, _, err := reader.ReadRune(); err == nil; input, _, err = reader.ReadRune() {
		if input == '\n' {
			break
		}

		num, err := strconv.Atoi(string(input))
		if err != nil {
			fmt.Println("Not a number:", err)
		}

		currId := -1
		if isFileId {
			currId = idSeq
			idSeq++
		}

		if currentNode2 == nil {
			disk2.first = &linkNode{fileId: currId, size: num, idx: idx2}
			currentNode2 = disk2.first
		} else {
			currentNode2.next = &linkNode{fileId: currId, prev: currentNode2, size: num, idx: idx2}
			currentNode2 = currentNode2.next
		}
		idx2++

		for i := 0; i < num; i++ {
			if currentNode == nil {
				disk.first = &linkNode{fileId: currId, idx: idx}
				currentNode = disk.first
			} else {
				currentNode.next = &linkNode{fileId: currId, prev: currentNode, idx: idx}
				currentNode = currentNode.next
			}
			idx++
		}

		isFileId = !isFileId
	}

	disk.last = currentNode
	disk2.last = currentNode2

	return &disk, &disk2
}

func Run() {
	disk, disk2 := parseInput()

	disk.reformat()
	fmt.Println("Part one:", disk.checkSum())

	disk2.reformatPartTwo()
	fmt.Println("Part two:", disk2.checkSumPartTwo())
}
