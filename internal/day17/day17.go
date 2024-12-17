package day17

import (
	"fmt"
	"reflect"
)

type device struct {
	registerA int
	registerB int
	registerC int

	ip           int
	instructions []int
	halt         bool

	outList []int

	debug bool
}

func (d *device) getComboValue(op int) int {
	switch op {
	case 4:
		return d.registerA
	case 5:
		return d.registerB
	case 6:
		return d.registerC
	default:
		return op
	}
}

func (d *device) adv(op int) {
	val := d.getComboValue(op)

	d.registerA = d.registerA / (1 << val)
}

func (d *device) bdv(op int) {
	val := d.getComboValue(op)

	d.registerB = d.registerA / (1 << val)
}

func (d *device) cdv(op int) {
	val := d.getComboValue(op)

	d.registerC = d.registerA / (1 << val)
}

func (d *device) bxl(op int) {
	d.registerB = d.registerB ^ op
}

func (d *device) bst(op int) {
	d.registerB = d.getComboValue(op) % 8
}

func (d *device) jnz(op int) {
	if d.registerA != 0 {
		d.ip = op
	}
}

func (d *device) bxc() {
	d.registerB = d.registerB ^ d.registerC
}

func (d *device) out(op int) {
	if d.debug {
		d.outList = append(d.outList, d.getComboValue(op)%8)

		if len(d.outList) > len(d.instructions) {
			d.halt = true
			return
		}

		return
	}

	fmt.Print(d.getComboValue(op)%8, ",")
}

func (d *device) execute() {
	if d.ip+1 > len(d.instructions) {
		d.halt = true
		return
	}

	opcode, op := d.instructions[d.ip], d.instructions[d.ip+1]

	d.ip += 2
	switch opcode {
	case 0:
		d.adv(op)
	case 1:
		d.bxl(op)
	case 2:
		d.bst(op)
	case 3:
		d.jnz(op)
	case 4:
		d.bxc()
	case 5:
		d.out(op)
	case 6:
		d.bdv(op)
	case 7:
		d.cdv(op)
	}
}

func findSolution(currA int, level int) bool {

	if level == 16 {
		return false
	}

	outs := []int{}
	for i := 0; i < 8; i++ {
		out := (currA << 3) | i
		dev := device{registerA: out, instructions: []int{}, debug: true}
		for !dev.halt {
			dev.execute()
		}
		if reflect.DeepEqual(dev.outList, dev.instructions[len(dev.instructions)-len(dev.outList):]) {
			outs = append(outs, out)

			if level == 15 {
				fmt.Println("Part two:", out)
				return true
			}
		}
	}

	for _, out := range outs {
		found := findSolution(out, level+1)
		if found {
			return true
		}
	}

	return false
}

func Run() {
	dev := device{registerA: 0, instructions: []int{}}

	fmt.Print("Part one: ")
	for !dev.halt {
		dev.execute()
	}
	fmt.Println()

	findSolution(0, 0)
}
