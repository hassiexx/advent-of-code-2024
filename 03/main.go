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
}

func part1() {
	input := readInput()

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

	fmt.Println("Result:", result)
}

func readInput() string {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(fmt.Errorf("failed to read input: %w", err))
	}

	return string(b)
}
