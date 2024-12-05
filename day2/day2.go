package day2

import (
	"fmt"
	"io"
	"math"
	"strings"

	"github.com/samber/lo"

	"github.com/a-bishop/aoc2024/utils"
)

func getLevels(report string) []int {
	return lo.Map(strings.Fields(report), func(field string, _ int) int {
		return utils.MustAtoi(field)
	})
}

func isDiffSafe(level1 int, level2 int, shouldBeDecreasing bool) bool {
	levelDiff := utils.AbsDiff(level1, level2)
	nowIsDecreasing := math.Signbit(float64(level2 - level1))
	if nowIsDecreasing != shouldBeDecreasing || levelDiff < 1 || levelDiff > 3 {
		return false
	}
	return true
}

func checkShouldBeDecreasing(levels []int) bool {
	return math.Signbit(float64(levels[1] - levels[0]))
}

func checkDiffs(levels []int) bool {
	fmt.Printf("Checking %v\n", levels)
	isDecr := checkShouldBeDecreasing(levels)
	for i := 1; i < len(levels); i++ {
		if !isDiffSafe(levels[i-1], levels[i], isDecr) {
			return false
		}
	}
	return true
}

func isReportSafe(levels []int, allowSkip bool) bool {
	fmt.Printf("Checking report %v allow skip %v\n", levels, allowSkip)
	if len(levels) < 2 {
		return true
	}
	isSafe := checkDiffs(levels)
	deleteIdx := 0
	for allowSkip && !isSafe && deleteIdx < len(levels) {
		// Successively remove the next element from the slice
		// TO FIX. NOT WORKING AS EXPECTED
		isSafe = checkDiffs(append(levels[:deleteIdx], levels[deleteIdx+1:]...))
		deleteIdx += 1
	}

	fmt.Printf("Safe: %v\n", isSafe)
	fmt.Println()

	return isSafe
}

func safeReportsMapping(reports [][]int) []bool {
	lookups := []bool{}
	for _, report := range reports {
		lookups = append(lookups, isReportSafe(report, false))
	}
	return lookups
}

func RedNosedReports(input io.Reader) (part1 int, part2 int) {
	reports := lo.Map(utils.GetLines(input), func(line string, _ int) []int {
		return getLevels(line)
	})

	safeReports := safeReportsMapping(reports)

	numSafeWithoutSkip := lo.Count(safeReports, true)

	numSafeWithSkip := 0
	for i, report := range reports {
		if safeReports[i] || isReportSafe(report, true) {
			numSafeWithSkip += 1
		}
	}

	return numSafeWithoutSkip, numSafeWithSkip
}
