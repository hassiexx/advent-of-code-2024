package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	result := calculateResult(readInput())

	fmt.Println("Result (part 1):", result)
}

func part2() {
	input := readInput()

	dontSpl := strings.Split(input, "don't()")

	var doInput string

	// Do is enabled by default so first split substring contains ops to do.
	doInput += dontSpl[0]
	dontSpl = dontSpl[1:]

	for _, dont := range dontSpl {
		index := strings.Index(dont, "do()")

		if index == -1 {
			continue
		}

		doInput += dont[index:]
	}

	result := calculateResult(doInput)

	fmt.Println("Result (part 2):", result)
}

func calculateResult(input string) int64 {
	r := regexp.MustCompile("mul\\(\\d+,\\d+\\)")
	matches := r.FindAllString(input, -1)

	var result int64

	for _, match := range matches {
		match = strings.NewReplacer("mul(", "", ")", "").Replace(match)
		spl := strings.Split(match, ",")

		ln, err := strconv.ParseInt(spl[0], 10, 32)
		if err != nil {
			panic(fmt.Errorf("failed to parse left number: %w", err))
		}

		rn, err := strconv.ParseInt(spl[1], 10, 32)
		if err != nil {
			panic(fmt.Errorf("failed to parse right number: %w", err))
		}

		result += ln * rn
	}

	return result
}

func readInput() string {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(fmt.Errorf("failed to read input: %w", err))
	}

	return string(b)
}
