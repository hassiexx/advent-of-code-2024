package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
}

type opcode uint8

const (
	opcodeAdv opcode = iota
	opcodeBxl
	opcodeBst
	opcodeJnz
	opcodeBxc
	opcodeOut
	opcodeBdv
	opcodeCdv
)

type program struct {
	Input []int64
	A     int64
	B     int64
	C     int64
	IP    int64
	Out   string
}

func part1() {
	p := readInput()

	for p.IP < int64(len(p.Input)) {
		executeInstruction(p)
	}

	fmt.Println("Output (part 1):", p.Out)
}

func executeInstruction(p *program) {
	jump := false

	switch p.Input[p.IP] {
	case int64(opcodeAdv):
		adv(p)
	case int64(opcodeBxl):
		bxl(p)
	case int64(opcodeBst):
		bst(p)
	case int64(opcodeJnz):
		jnz(p)
		jump = true
	case int64(opcodeBxc):
		bxc(p)
	case int64(opcodeOut):
		out(p)
	case int64(opcodeBdv):
		bdv(p)
	case int64(opcodeCdv):
		cdv(p)
	}

	if !jump {
		p.IP += 2
	}
}

func adv(p *program) {
	p.A = dv(p.A, comboOprValue(p))
}

func bxl(p *program) {
	p.B ^= opr(p)
}

func bst(p *program) {
	p.B = comboOprValue(p) % 8
}

func bdv(p *program) {
	p.B = dv(p.A, comboOprValue(p))
}

func cdv(p *program) {
	p.C = dv(p.A, comboOprValue(p))
}

func jnz(p *program) {
	if p.A == 0 {
		p.IP += 2
		return
	}

	p.IP = p.Input[p.IP+1]

	if p.Input[p.IP] == int64(opcodeJnz) {
		return
	}

	executeInstruction(p)
}

func bxc(p *program) {
	p.B ^= p.C
}

func out(p *program) {
	v := comboOprValue(p) % 8

	if p.Out == "" {
		p.Out = strconv.Itoa(int(v))
	} else {
		p.Out += "," + strconv.Itoa(int(v))
	}
}

func dv(n int64, d int64) int64 {
	return int64(float64(n) / math.Pow(2, float64(d)))
}

func comboOprValue(p *program) int64 {
	opr := p.Input[p.IP+1]

	switch opr {
	case 0, 1, 2, 3:
		return opr
	case 4:
		return p.A
	case 5:
		return p.B
	case 6:
		return p.C
	case 7:
		panic("Opcode 7 not handled")
	default:
		panic("Unknown opcode")
	}
}

func opr(p *program) int64 {
	return p.Input[p.IP+1]
}

func readInput() *program {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(fmt.Errorf("failed to read file: %w", err))
	}

	defer f.Close()

	sc := bufio.NewScanner(f)
	registers := make([]int64, 3)

	for i := range 3 {
		sc.Scan()

		v, err := strconv.ParseInt(strings.Split(sc.Text(), ": ")[1], 10, 64)
		if err != nil {
			panic(fmt.Errorf("failed to parse register value: %w", err))
		}

		registers[i] = v
	}

	sc.Scan()
	sc.Scan()

	inputStr := strings.Split(strings.Split(sc.Text(), ": ")[1], ",")
	input := make([]int64, len(inputStr))

	for i, v := range inputStr {
		op, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(fmt.Errorf("failed to parse opcode/operand: %w", err))
		}

		input[i] = op
	}

	return &program{
		Input: input,
		A:     registers[0],
		B:     registers[1],
		C:     registers[2],
	}
}
