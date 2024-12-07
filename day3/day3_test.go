package day3

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sampleInput1 = strings.NewReader(`xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`)

var sampleInput2 = strings.NewReader(`xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`)

func TestMullItOver(t *testing.T) {
	part1, _ := MullItOver(sampleInput1)
	assert.Equal(t, 161, part1)

	_, part2 := MullItOver(sampleInput2)
	assert.Equal(t, 48, part2)

}
