package aoc2020

import (
	"github.com/tlcowling/adventutils"
	"sort"
	"testing"
)

func TestDay10(t *testing.T) {
	lines := adventutils.ReadInputAsInts("./inputs/day10.txt")
	sort.Ints(lines)
	maxDeviceVolts := lines[len(lines)-1] + 3
	lines = append(lines, maxDeviceVolts)

	joltMap := make(map[int]int)
	lastJolt := 0

	for _, thisJolt := range lines {
		joltMap[thisJolt-lastJolt]++
		lastJolt = thisJolt
	}

	t.Log(joltMap[1] * joltMap[3])

	t.Log(countCombos(lines))
}

func countCombos(jolts []int) int {
	counts := map[int]int{0: 1}

	for _, jolt := range jolts {
		counts[jolt] = counts[jolt-3] + counts[jolt-2] + counts[jolt-1]
	}
	return counts[jolts[len(jolts)-1]]
}
