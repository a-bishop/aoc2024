package day1

import (
	"bufio"
	"io"
	"slices"
	"strings"

	"github.com/a-bishop/aoc2024/utils"
)

func getSortedLocationLists(input io.Reader) ([]int, []int) {
	var leftList, rightList []int
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		spl := strings.Fields(scanner.Text())

		num1 := utils.MustAtoi(spl[0])
		num2 := utils.MustAtoi(spl[1])

		leftList = append(leftList, num1)
		rightList = append(rightList, num2)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	return leftList, rightList
}

func getDistanceAndRightListTally(leftList, rightList []int) (int, map[int]int) {
	distance := 0
	var rightListTally = map[int]int{}
	for idx := range leftList {
		distance += utils.AbsDiff(leftList[idx], rightList[idx])
		rightListTally[rightList[idx]] += 1
	}

	return distance, rightListTally
}

func getSimScore(leftList []int, rightListTally map[int]int) int {
	simScore := 0
	for _, num := range leftList {
		simScore += num * rightListTally[num]
	}

	return simScore
}

func HistorianHysteria(input io.Reader) (part1 int, part2 int) {
	leftList, rightList := getSortedLocationLists(input)

	distance, rightListTally := getDistanceAndRightListTally(leftList, rightList)

	simScore := getSimScore(leftList, rightListTally)

	return distance, simScore
}
