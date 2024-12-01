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
	file, err := os.Open("input/01-input.txt")
	if err != nil {
		panic(err)
	}

	left := make([]int32, 0, 1000)
	right := make([]int32, 0, 1000)

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		t := sc.Text()
		spl := strings.Split(t, "   ")

		l, err := strconv.ParseInt(spl[0], 10, 32)
		if err != nil {
			panic(fmt.Errorf("failed to parse left number: %w", err))
		}

		r, err := strconv.ParseInt(spl[1], 10, 32)
		if err != nil {
			panic(fmt.Errorf("failed to parse right number: %w", err))
		}

		left = append(left, int32(l))
		right = append(right, int32(r))
	}

	if err := sc.Err(); err != nil {
		panic(fmt.Errorf("failed to scan file: %w", err))
	}

	slices.Sort(left)
	slices.Sort(right)

	var distance int64
	for i := 0; i < len(left); i++ {
		diff := left[i] - right[i]
		distance += int64(math.Abs(float64(diff)))
	}

	fmt.Println("Distance: " + strconv.Itoa(int(distance)))
}
