package day4

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sampleInput = strings.NewReader(`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`)

func TestCeresSearch(t *testing.T) {
	part1, part2 := CeresSearch(sampleInput)
	assert.Equal(t, 18, part1)
	assert.Equal(t, 9, part2)
}
