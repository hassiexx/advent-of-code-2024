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
	left, right := readInput()

	slices.Sort(left)
	slices.Sort(right)

	var distance int64
	for i := 0; i < len(left); i++ {
		diff := left[i] - right[i]
		distance += int64(math.Abs(float64(diff)))
	}

	fmt.Println("Distance: " + strconv.Itoa(int(distance)))
}

func part2() {
	left, right := readInput()

	appearances := make(map[int32]int32)

	for _, v := range left {
		appearances[v] = 0
	}

	for _, v := range right {
		_, ok := appearances[v]
		if !ok {
			continue
		}

		appearances[v] += 1
	}

	var similarity int64
	for _, v := range left {
		similarity += int64(v * appearances[v])
	}

	fmt.Println("Similarity: " + strconv.Itoa(int(similarity)))
}

func readInput() (left []int32, right []int32) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	left = make([]int32, 0, 1000)
	right = make([]int32, 0, 1000)

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		t := sc.Text()
		spl := strings.Split(t, "   ")

		ln, err := strconv.ParseInt(spl[0], 10, 32)
		if err != nil {
			panic(fmt.Errorf("failed to parse left number: %w", err))
		}

		rn, err := strconv.ParseInt(spl[1], 10, 32)
		if err != nil {
			panic(fmt.Errorf("failed to parse right number: %w", err))
		}

		left = append(left, int32(ln))
		right = append(right, int32(rn))
	}

	if err := sc.Err(); err != nil {
		panic(fmt.Errorf("failed to scan file: %w", err))
	}

	return left, right
}
