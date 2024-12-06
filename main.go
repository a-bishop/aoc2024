package main

import (
	"fmt"
	"io"
	"os"

	"github.com/a-bishop/aoc2024/day1"
	"github.com/a-bishop/aoc2024/day2"
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
}
