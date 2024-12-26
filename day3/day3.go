package day3

import (
	"io"
	"regexp"
	"strings"

	"github.com/a-bishop/aoc2024/utils"
)

func MullItOver(input io.Reader) (part1 int, part2 int) {
	corruptedMem := strings.Join(utils.GetLines(input), "\n")

	re := regexp.MustCompile(`don\'t\(\)|do\(\)|mul\((\d+),(\d+)\)`)

	total := 0
	totalWhileEnabled := 0

	isEnabled := true
	for _, match := range re.FindAllStringSubmatch(corruptedMem, -1) {
		if match[0] == "don't()" {
			isEnabled = false
		} else if match[0] == "do()" {
			isEnabled = true
		} else {
			sum := utils.MustAtoi(match[1]) * utils.MustAtoi(match[2])
			total += sum
			if isEnabled {
				totalWhileEnabled += sum
			}
		}
	}

	return total, totalWhileEnabled
}
