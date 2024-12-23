package day23

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/DanJsef/AoC2024/internal/utils"
)

type node struct {
	value     string
	candidate bool
	edges     []string
}

type graph struct {
	nodes   map[string]*node
	cycles  map[[3]string]bool
	cliques map[string]bool
}

func (g graph) dfs(n *node, path []string, start string) {
	path = append(path, n.value)
	for _, edge := range n.edges {
		if len(path) < 3 {
			g.dfs(g.nodes[edge], path, start)
			continue
		}
		if edge == start {
			cycle := [3]string{path[0], path[1], path[2]}
			sort.Strings(cycle[:])
			g.cycles[cycle] = true
		}
	}
}

func (g graph) searchConnections() {
	for _, node := range g.nodes {
		if !node.candidate {
			continue
		}

		g.dfs(node, []string{}, node.value)
	}
}

func (g graph) findCliques(R, P, X []string) {
	if len(P) == 0 && len(X) == 0 {
		clique := append([]string{}, R...)
		sort.Strings(clique)
		g.cliques[strings.Join(clique, ",")] = true
	}

	for _, n := range append([]string{}, P...) {
		g.findCliques(append(R, n), utils.IntersectSlices(P, g.nodes[n].edges), utils.IntersectSlices(X, g.nodes[n].edges))

		P = utils.RemoveFromSlice(P, n)
		X = append(X, n)
	}
}

func (g graph) findLanParty() string {
	g.findCliques([]string{}, utils.GetMapKeys(g.nodes), []string{})

	largest := ""
	for clique := range g.cliques {
		if len(clique) > len(largest) {
			largest = clique
		}
	}

	return largest
}

func parseInput() *graph {
	file, err := os.Open("./inputs/day23.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	g := &graph{nodes: map[string]*node{}, cycles: map[[3]string]bool{}, cliques: map[string]bool{}}

	for line, err := reader.ReadString('\n'); err == nil; line, err = reader.ReadString('\n') {
		line = line[:len(line)-1]

		nodeA, nodeB := line[:2], line[3:]

		if _, ok := g.nodes[nodeA]; !ok {
			g.nodes[nodeA] = &node{value: nodeA, candidate: nodeA[0] == 't'}
		}

		if _, ok := g.nodes[nodeB]; !ok {
			g.nodes[nodeB] = &node{value: nodeB, candidate: nodeB[0] == 't'}
		}

		g.nodes[nodeA].edges = append(g.nodes[nodeA].edges, nodeB)
		g.nodes[nodeB].edges = append(g.nodes[nodeB].edges, nodeA)
	}

	return g
}

func Run() {
	g := parseInput()

	g.searchConnections()
	fmt.Println("Part one:", len(g.cycles))

	fmt.Println("Part two:", g.findLanParty())
}
