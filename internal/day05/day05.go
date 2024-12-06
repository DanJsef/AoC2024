package day05

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	value         int
	visited       bool
	done          bool
	adjacentNodes []int
}

func (n *node) visit(g *ruleGraph, acc *[]int) {
	if n.done {
		return
	}

	if n.visited {
		fmt.Println("Well, the assumption that the graph is a DAG was wrong xD")
		return
	}

	n.visited = true

	for _, adj := range n.adjacentNodes {
		if adjNode, ok := (*g)[adj]; !ok {
			*acc = append(*acc, adj)
			(*g)[adj] = &node{value: adj, done: true}
			continue
		} else {
			adjNode.visit(g, acc)
		}
	}

	*acc = append(*acc, n.value)
	n.done = true
}

func makePageFilter(page []int) map[int]bool {
	filter := make(map[int]bool)
	for _, val := range page {
		filter[val] = true
	}
	return filter
}

type ruleGraph map[int]*node

func (g *ruleGraph) filterGraph(filter map[int]bool) ruleGraph {

	filteredGraph := make(ruleGraph)

	for _, n := range *g {
		if _, ok := filter[n.value]; !ok {
			continue
		}

		newAdj := []int{}
		for _, adj := range n.adjacentNodes {
			if _, ok := filter[adj]; ok {
				newAdj = append(newAdj, adj)
			}
		}
		filteredGraph[n.value] = &node{value: n.value, adjacentNodes: newAdj}
	}

	return filteredGraph
}

func (g ruleGraph) getOrdering(page []int) pageOrdering {

	filter := makePageFilter(page)

	fg := g.filterGraph(filter)

	pageOrdering := make(pageOrdering)

	orderAcc := []int{}

	for _, node := range fg {
		node.visit(&fg, &orderAcc)
	}

	for idx, val := range orderAcc {
		pageOrdering[val] = len(orderAcc) - idx - 1
	}

	return pageOrdering
}

type pageOrdering map[int]int

func (p *pageOrdering) checkPage(page []int) (int, bool) {

	pos := -1

	for _, val := range page {
		if order, ok := (*p)[val]; ok {
			if pos < order {
				pos = order
				continue
			}
			return 0, false
		}
	}

	return page[len(page)/2], true
}

type pageList [][]int

func parseInput() (ruleGraph, pageList) {
	file, err := os.Open("./inputs/day05.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	graph := make(ruleGraph)

	for line, err := reader.ReadString('\n'); err == nil && line != "\n"; line, err = reader.ReadString('\n') {
		line = strings.TrimSpace(line)
		split := strings.Split(line, "|")

		value, _ := strconv.Atoi(split[0])
		adjecent, _ := strconv.Atoi(split[1])

		valueNode, ok := graph[value]

		if ok {
			valueNode.adjacentNodes = append(valueNode.adjacentNodes, adjecent)
			continue
		}

		graph[value] = &node{value: value, adjacentNodes: []int{adjecent}}
	}

	pages := pageList{}

	for line, err := reader.ReadString('\n'); err == nil; line, err = reader.ReadString('\n') {
		line = strings.TrimSpace(line)

		split := strings.Split(line, ",")

		var page []int

		for _, val := range split {
			parsed, _ := strconv.Atoi(val)
			page = append(page, parsed)
		}

		pages = append(pages, page)
	}

	return graph, pages
}

func Run() {
	graph, pages := parseInput()

	partOne := 0
	partTwo := 0

	for _, page := range pages {
		ordering := graph.getOrdering(page)
		if val, ok := ordering.checkPage(page); ok {
			partOne += val
		} else {
			for value, order := range ordering {
				if order == len(ordering)/2 {
					partTwo += value
				}
			}
		}
	}

	fmt.Println("Part one:", partOne)
	fmt.Println("Part two:", partTwo)
}
