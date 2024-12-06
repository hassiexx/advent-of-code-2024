package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	rules, updates := readInput()

	var result int32

	for _, update := range updates {
		if !slices.IsSortedFunc(update, sortFunc(rules)) {
			continue
		}

		mid := update[int32(math.Round(float64(len(update)/2)))]
		midI, err := strconv.ParseInt(mid, 10, 32)
		if err != nil {
			panic(fmt.Errorf("failed to parse page number: %w", err))
		}

		result += int32(midI)
	}

	fmt.Println("Result (part 1):", result)
}

func part2() {
	rules, updates := readInput()

	var result int32

	for _, update := range updates {
		if slices.IsSortedFunc(update, sortFunc(rules)) {
			continue
		}

		slices.SortStableFunc(update, sortFunc(rules))

		mid := update[len(update)/2]
		midI, err := strconv.ParseInt(mid, 10, 32)
		if err != nil {
			panic(fmt.Errorf("failed to parse page number: %w", err))
		}

		result += int32(midI)
	}

	fmt.Println("Result (part 2)", result)
}

func sortFunc(rules map[string]struct{}) func(a string, b string) int {
	return func(a, b string) int {
		k := a + "|" + b

		if _, ok := rules[k]; ok {
			return -1
		}

		return 1
	}
}

// Alternative solution to part 1 not using built in slice sort func.
func part1Alternative() {
	rules, updates := readInput()

	var result int32

	for _, update := range updates {
		skip := false

	outer:
		for i := 0; i < len(update)-1; i++ {
			for j := i + 1; j < len(update); j++ {
				rule := update[i] + "|" + update[j]

				if _, ok := rules[rule]; !ok {
					skip = true
					break outer
				}
			}
		}

		if skip {
			continue
		}

		mid := update[len(update)/2]
		midI, err := strconv.ParseInt(mid, 10, 32)
		if err != nil {
			panic(fmt.Errorf("failed to parse page number: %w", err))
		}

		result += int32(midI)
	}

	fmt.Println("Result (part 1):", result)
}

func readInput() (rules map[string]struct{}, updates [][]string) {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(fmt.Errorf("failed to open input file: %w", err))
	}

	defer f.Close()

	sc := bufio.NewScanner(f)

	rules = make(map[string]struct{})
	for sc.Scan() {
		if sc.Text() == "" {
			break
		}

		rules[sc.Text()] = struct{}{}
	}

	updates = make([][]string, 0)
	for sc.Scan() {
		updates = append(updates, strings.Split(sc.Text(), ","))
	}

	if sc.Err() != nil {
		panic(fmt.Errorf("failed to scan input: %w", err))
	}

	return rules, updates
}
