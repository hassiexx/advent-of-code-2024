package main

import (
	"bufio"
	"fmt"
	"os"
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
