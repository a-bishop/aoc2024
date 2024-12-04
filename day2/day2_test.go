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

func TestDay2(t *testing.T) {
	part1, _ := Day2(sampleInput)

	assert.Equal(t, 2, part1)
}
