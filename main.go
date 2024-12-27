package main

import (
	"fmt"
	"io"
	"os"

	"github.com/a-bishop/aoc2024/day1"
	"github.com/a-bishop/aoc2024/day2"
	"github.com/a-bishop/aoc2024/day3"
	"github.com/a-bishop/aoc2024/day4"
	"github.com/a-bishop/aoc2024/day5"
	"github.com/a-bishop/aoc2024/utils"
)

func runDay(dayNum int, file *os.File, f func(input io.Reader) (int, int)) {
	defer file.Close()
	part1, part2 := f(file)
	fmt.Printf("Day %d, Part 1: %d\n", dayNum, part1)
	fmt.Printf("Day %d, Part 2: %d\n", dayNum, part2)
}

func main() {
	runDay(1, utils.GetFile("inputs/day1.txt"), day1.HistorianHysteria)
	fmt.Println()

	runDay(2, utils.GetFile("inputs/day2.txt"), day2.RedNosedReports)
	fmt.Println()

	runDay(3, utils.GetFile("inputs/day3.txt"), day3.MullItOver)
	fmt.Println()

	runDay(4, utils.GetFile("inputs/day4.txt"), day4.CeresSearch)
	fmt.Println()

	runDay(5, utils.GetFile("inputs/day5.txt"), day5.PrintQueue)
	fmt.Println()
}
