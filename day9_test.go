package aoc2020

import (
	"github.com/tlcowling/adventutils"
	"testing"
)

func TestDay9(t *testing.T) {
	lines := adventutils.ReadInputAsInts("./inputs/day9.txt")
	preambleSize := 25
	preamble := lines[0:preambleSize]
	rest := lines[preambleSize:]
	lastValid := []int{}

	for i := 0; i < preambleSize; i++ {
		lastValid = append(lastValid, preamble[i])
	}

	var firstInvalidNumber int
	for _, p := range rest {
		if !arrayContainsSumToN(lastValid, p) {
			firstInvalidNumber = p
			t.Log(firstInvalidNumber)
			break
		}
		lastValid = append(lastValid[1:], p)
	}

	target := firstInvalidNumber
	for i := 0; i < len(lines)-1; i++ {
		sumNums := []int{lines[i]}
		sumFromI := lines[i]
		for j := i + 1; j < len(lines); j++ {
			sumFromI += lines[j]
			sumNums = append(sumNums, lines[j])
			if sumFromI == target {
				min, max := adventutils.ArrayMinMax(sumNums)
				t.Log(min + max)
				return
			}
		}
	}
}

func arrayContainsSumToN(arr []int, n int) bool {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i == j {
				continue
			}
			if arr[i]+arr[j] == n {
				return true
			}
		}
	}
	return false
}
