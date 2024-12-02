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
	file, err := os.Open("../inputs/day1.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	var left, right []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		spl := strings.Split(scanner.Text(), "   ")

		num1, _ := strconv.Atoi(spl[0])
		num2, _ := strconv.Atoi(spl[1])

		left = append(left, num1)
		right = append(right, num2)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	slices.Sort(left)
	slices.Sort(right)

	distance := 0
	var rCounts = map[int]int{}
	for idx := range left {
		distance += int(math.Abs(float64(right[idx] - left[idx])))
		rCounts[right[idx]] += 1
	}

	fmt.Printf("Part 1: %v\n", distance)

	simScore := 0
	for _, num := range left {
		simScore += num * rCounts[num]
	}
	fmt.Printf("Part 2: %v\n", simScore)
}
