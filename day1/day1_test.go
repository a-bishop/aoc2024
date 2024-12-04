package day1

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sampleInput = strings.NewReader(
	`3   4
4   3
2   5
1   3
3   9
3   3`)

func TestDay1(t *testing.T) {
	part1, part2 := Day1(sampleInput)

	assert.Equal(t, 11, part1)
	assert.Equal(t, 31, part2)
}
