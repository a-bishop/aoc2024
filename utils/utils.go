package utils

import (
	"bufio"
	"io"
	"math"
	"os"
	"strconv"
)

func GetFile(filePath string) *os.File {
	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	return file
}

func GetLines(input io.Reader) []string {
	scanner := bufio.NewScanner(input)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func GetLinesWithSplit(input io.Reader, splitOn string) ([]string, []string) {
	scanner := bufio.NewScanner(input)

	lines1 := []string{}
	lines2 := []string{}
	hasSplit := false

	for scanner.Scan() {
		txt := scanner.Text()
		if txt == splitOn {
			hasSplit = true
		} else if hasSplit {
			lines2 = append(lines2, txt)
		} else {
			lines1 = append(lines1, txt)
		}
	}

	return lines1, lines2
}

func AbsDiff(num1, num2 int) int {
	return int(math.Abs(float64(num1 - num2)))
}

func MustAtoi(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}

func RemoveAtIndex[T any](s []T, index int) []T {
	ns := make([]T, 0, len(s))
	ns = append(ns, s[:index]...)
	return append(ns, s[index+1:]...)
}
