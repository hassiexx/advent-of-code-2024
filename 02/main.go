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
	reports := readInput()

	var count int32

	for _, report := range reports {
		if isSafe(report) {
			count++
		}
	}

	fmt.Println("Safe reports (part 1):", count)
}

func part2() {
	reports := readInput()

	var count int32

	for _, report := range reports {
		if isSafe(report) {
			count++
			continue
		}

		for i := 0; i < len(report); i++ {
			dampenedReport := slices.Clone(report)
			dampenedReport = append(dampenedReport[:i], dampenedReport[i+1:]...)

			if isSafe(dampenedReport) {
				count++
				break
			}
		}
	}

	fmt.Println("Safe reports (part 2):", count)
}

func isSafe(report []int8) bool {
	ascending := report[1] > report[0]

	for i := 1; i < len(report); i++ {
		diff := math.Abs(float64(report[i]) - float64(report[i-1]))
		if diff < 1 || diff > 3 {
			return false
		}

		if (report[i] > report[i-1] && ascending) || (report[i] < report[i-1] && !ascending) {
			continue
		}

		return false
	}

	return true
}

func readInput() [][]int8 {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(fmt.Errorf("failed to open file: %w", err))
	}

	defer f.Close()

	reports := make([][]int8, 0, 1000)

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		t := sc.Text()
		spl := strings.Split(t, " ")

		report := make([]int8, len(spl))

		for i := 0; i < len(spl); i++ {
			l, err := strconv.ParseInt(spl[i], 10, 8)
			if err != nil {
				panic(fmt.Errorf("failed to parse level in report: %w", err))
			}

			report[i] = int8(l)
		}

		reports = append(reports, report)
	}

	if sc.Err() != nil {
		panic(fmt.Errorf("failed to scan reports: %w", err))
	}

	return reports
}
