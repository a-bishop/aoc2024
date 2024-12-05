package day2

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sampleInput = strings.NewReader(
	`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`)

func TestRedNosedReports(t *testing.T) {
	part1, part2 := RedNosedReports(sampleInput)

	assert.Equal(t, 2, part1)
	assert.Equal(t, 4, part2)
}
