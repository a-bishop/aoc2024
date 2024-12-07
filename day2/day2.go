package day2

import (
	"io"
	"strings"

	"github.com/samber/lo"

	"github.com/a-bishop/aoc2024/utils"
)

func getLevels(report string) []int {
	return lo.Map(strings.Fields(report), func(field string, _ int) int {
		return utils.MustAtoi(field)
	})
}

func getReports(input io.Reader) [][]int {
	return lo.Map(utils.GetLines(input), func(line string, _ int) []int {
		return getLevels(line)
	})
}

func isDiffSafe(level1 int, level2 int, shouldBeDecreasing bool) bool {
	levelDiff := utils.AbsDiff(level2, level1)
	nowIsDecreasing := level2 > level1
	if nowIsDecreasing != shouldBeDecreasing || levelDiff < 1 || levelDiff > 3 {
		return false
	}
	return true
}

func checkDiffs(levels []int) bool {
	isDecr := levels[1] > levels[0]
	for i := 1; i < len(levels); i++ {
		if !isDiffSafe(levels[i-1], levels[i], isDecr) {
			return false
		}
	}
	return true
}

func isReportSafe(levels []int, allowSkipOne bool) bool {
	if len(levels) < 2 {
		return true
	}
	isSafe := checkDiffs(levels)
	if !isSafe && allowSkipOne {
		deleteIdx := 0
		for !isSafe && deleteIdx < len(levels) {
			// Successively remove the next element from the slice until safe or options exhausted
			isSafe = checkDiffs(utils.RemoveAtIndex(levels, deleteIdx))
			deleteIdx += 1
		}
	}
	return isSafe
}

func RedNosedReports(input io.Reader) (part1 int, part2 int) {
	reports := getReports(input)

	numSafe := 0
	numProblemDampened := 0
	for _, report := range reports {
		if isReportSafe(report, false) {
			numSafe += 1
		} else if isReportSafe(report, true) {
			numProblemDampened += 1
		}
	}

	return numSafe, numSafe + numProblemDampened
}
