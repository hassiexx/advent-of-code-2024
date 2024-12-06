package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	part1()
}

type direction uint32

const (
	directionNorth direction = 0
	directionEast  direction = 90
	directionSouth direction = 180
	directionWest  direction = 270
)

type pos struct {
	x int32
	y int32
}

func part1() {
	m, p := readInput()
	d := directionNorth
	guarding := true
	var visited uint32 = 1 // The starting position counts as visited.

	for guarding {
		var distinctPos bool
		p, d, distinctPos, guarding = traverse(m, p, d)

		if distinctPos {
			visited += 1
		}
	}

	fmt.Println("Distinct positions (part 1):", visited)
}

func traverse(m [][]rune, p pos, d direction) (newPos pos, newDirection direction, distinctPos bool, guarding bool) {
	newPos = p

	switch d {
	case directionNorth:
		newPos.y -= 1
	case directionEast:
		newPos.x += 1
	case directionSouth:
		newPos.y += 1
	case directionWest:
		newPos.x -= 1
	}

	// Guard gone off map
	if newPos.x < 0 || newPos.x == int32(len(m[0])) || newPos.y < 0 || newPos.y == int32(len(m)) {
		return newPos, d, false, false
	}

	// Guard hit an obstacle.
	if m[newPos.y][newPos.x] == '#' {
		// Rotate 90 deg clockwise and traverse again.
		if d == directionWest {
			d = directionNorth
		} else {
			d += 90
		}

		return traverse(m, p, d)
	}

	// Mark as visited if distinct.
	if m[newPos.y][newPos.x] == '.' {
		m[newPos.y][newPos.x] = 'X'
		return newPos, d, true, true
	}

	return newPos, d, false, true
}

func readInput() (m [][]rune, sp pos) {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(fmt.Errorf("failed to read file: %w", err))
	}

	defer f.Close()

	m = make([][]rune, 0)

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		m = append(m, []rune(sc.Text()))
	}

	if err := sc.Err(); err != nil {
		panic(fmt.Errorf("failed to scan input: %w", err))
	}

outer:
	// Mark starting position as visited.
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			if m[y][x] == '^' {
				m[y][x] = 'X'
				sp.x = int32(x)
				sp.y = int32(y)
				break outer
			}
		}
	}

	return m, sp
}
