package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
}

func part1() {
	stones := readInput()

	const passes = 25

	for range passes {
		var newStones int

		for i := 0; i < len(stones)-newStones; i++ {
			switch {
			case stones[i] == 0:
				stones[i] = 1
			case hasEvenDigits(stones[i]):
				a, b := splitStone(stones[i])
				stones[i] = a
				stones = append(stones, b)
				newStones++
			default:
				stones[i] = stones[i] * 2024
			}
		}
	}

	fmt.Println("Number of stones (part 1):", len(stones))
}

func hasEvenDigits(i int64) bool {
	var digits int64

	for i != 0 {
		digits++
		i /= 10
	}

	return digits%2 == 0
}

func splitStone(i int64) (a, b int64) {
	s := strconv.Itoa(int(i))

	as := s[:len(s)/2]
	bs := s[len(s)/2:]

	a, _ = strconv.ParseInt(as, 10, 64)
	b, _ = strconv.ParseInt(bs, 10, 64)

	return a, b
}

func readInput() []int64 {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(fmt.Errorf("failed to read file: %w", err))
	}

	spl := strings.Split(string(b), " ")

	stones := make([]int64, len(spl))
	for i, s := range spl {
		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(fmt.Errorf("failed to parse number: %w", err))
		}

		stones[i] = n
	}

	return stones
}
