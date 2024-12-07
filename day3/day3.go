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
	totalEnabled := 0
	enabled := true
	matches := re.FindAllStringSubmatch(corruptedMem, -1)
	for _, match := range matches {
		if match[0] == "don't()" {
			enabled = false
		} else if match[0] == "do()" {
			enabled = true
		} else {
			sum := utils.MustAtoi(match[1]) * utils.MustAtoi(match[2])
			total += sum
			if enabled {
				totalEnabled += sum
			}
		}
	}

	return total, totalEnabled
}
