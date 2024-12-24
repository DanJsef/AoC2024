package day24

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/DanJsef/AoC2024/internal/utils"
)

type gate struct {
	wireInA  string
	operator string
	wireInB  string
	wireOut  string
}

type device struct {
	wireValues map[string]int
	gates      []*gate
	outValues  map[int]int
}

func (d *device) simulate() {
	for len(d.gates) > 0 {
		for i, g := range d.gates {
			if _, ok := d.wireValues[g.wireInA]; !ok {
				continue
			}
			if _, ok := d.wireValues[g.wireInB]; !ok {
				continue
			}
			switch g.operator {
			case "AND":
				d.wireValues[g.wireOut] = d.wireValues[g.wireInA] & d.wireValues[g.wireInB]
			case "OR":
				d.wireValues[g.wireOut] = d.wireValues[g.wireInA] | d.wireValues[g.wireInB]
			case "XOR":
				d.wireValues[g.wireOut] = d.wireValues[g.wireInA] ^ d.wireValues[g.wireInB]
			}
			d.gates = append(d.gates[:i], d.gates[i+1:]...)

			if g.wireOut[0] == 'z' {
				d.outValues[utils.ExtractNumbers(g.wireOut)[0]] = d.wireValues[g.wireOut]
			}

			break
		}
	}
}

func (d *device) output() int {
	out := 0

	for k, v := range d.outValues {
		out += v << k
	}

	return out
}

func parseInput() *device {
	file, err := os.Open("./inputs/day24.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	d := &device{wireValues: map[string]int{}, outValues: map[int]int{}}
	init := true

	for line, err := reader.ReadString('\n'); err == nil; line, err = reader.ReadString('\n') {
		if line == "\n" {
			init = false
			continue
		}

		line = line[:len(line)-1]

		if init {
			parts := strings.Split(line, ":")
			d.wireValues[parts[0]] = utils.ExtractNumbers(parts[1])[0]
			continue
		}

		parts := strings.Fields(line)

		d.gates = append(d.gates, &gate{wireInA: parts[0], operator: parts[1], wireInB: parts[2], wireOut: parts[4]})
	}

	return d
}

func sortStrings(arr [8]string) string {
	slices.Sort(arr[:])
	return strings.Join(arr[:], ",")
}

func Run() {
	d := parseInput()

	d.simulate()

	fmt.Println("Part one:", d.output())
	// Part two was solved by visualizing the output (transforming the input file using Vim macro into DOT format) and manually correcting the wrong wires one by one.
	fmt.Println("Part two:", sortStrings([8]string{"kdf", "z23", "rpp", "z39", "fdv", "dbp", "ckj", "z15"}))
}
