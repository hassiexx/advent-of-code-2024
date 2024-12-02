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

func part1() {
	reports := readInput()

	var count int32

	for _, report := range reports {
		ascending := report[1] > report[0]
		safe := true

		for i := 1; i < len(report); i++ {
			diff := math.Abs(float64(report[i]) - float64(report[i-1]))
			if diff < 1 || diff > 3 {
				safe = false
				break
			}

			if (report[i] > report[i-1] && ascending) || (report[i] < report[i-1] && !ascending) {
				continue
			}

			safe = false
			break
		}

		if safe {
			count++
		}
	}

	fmt.Println("Safe reports:", count)
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
