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
