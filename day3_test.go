package aoc2020

import (
	"strings"
	"testing"
)

func TestDay3(t *testing.T) {
	lines := ReadInputAsLines(3)

	a1 := slopeCount(lines, 1, 1)
	a2 := slopeCount(lines, 3, 1)
	a3 := slopeCount(lines, 5, 1)
	a4 := slopeCount(lines, 7, 1)
	a5 := slopeCount(lines, 1, 2)
	t.Log(a2)
	t.Log(a1 * a2 * a3 * a4 * a5)
}

func slopeCount(lines []string, right, down int) int {
	i := 0
	j := 0
	count := 0
	height := len(lines)
	for i < height-1 {
		j += right
		i += down
		line := lines[i]
		points := strings.Split(line, "")
		jpos := j % len(line)
		contact := points[jpos]
		if contact == "#" {
			count++
		}
	}
	return count
}
