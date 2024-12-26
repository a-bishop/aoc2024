package day2

import (
	"io"
	"strings"

	"github.com/samber/lo"

	"github.com/a-bishop/aoc2024/utils"
)

type Levels []int

type Reports []Levels

func getLevels(report string) Levels {
	return lo.Map(strings.Fields(report), func(field string, _ int) int {
		return utils.MustAtoi(field)
	})
}

func getReports(input io.Reader) Reports {
	return lo.Map(utils.GetLines(input), func(line string, _ int) Levels {
		return getLevels(line)
	})
}

func isDiffSafe(level1 int, level2 int, shouldBeIncreasing bool) bool {
	levelDiff := utils.AbsDiff(level2, level1)
	nowIsIncreasing := level1 > level2
	if nowIsIncreasing != shouldBeIncreasing || levelDiff < 1 || levelDiff > 3 {
		return false
	}
	return true
}

func checkDiffs(levels Levels) bool {
	isIncr := levels[0] > levels[1]
	for i := 1; i < len(levels); i++ {
		if !isDiffSafe(levels[i-1], levels[i], isIncr) {
			return false
		}
	}
	return true
}

func isReportSafe(levels Levels, allowSkipOne bool) bool {
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
