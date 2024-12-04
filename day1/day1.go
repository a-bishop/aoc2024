package day1

import (
	"bufio"
	"io"
	"math"
	"slices"
	"strconv"
	"strings"
)

func Day1(input io.Reader) (int, int) {
	var left, right []int
	scanner := bufio.NewScanner(input)
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

	simScore := 0
	for _, num := range left {
		simScore += num * rCounts[num]
	}

	return distance, simScore
}
