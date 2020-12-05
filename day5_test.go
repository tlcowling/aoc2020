package aoc2020

import (
	"github.com/tlcowling/adventutils"
	"strings"
	"testing"
)

func TestDay5(t *testing.T) {
	lines := adventutils.ReadInputAsLines("./inputs/day5.txt")

	seatMap := make(map[int]bool)
	maxSeatId := 0

	for _, line := range lines {
		rowMin, rowMax, colMin, colMax := 0, 127, 0, 7

		tokens := strings.Split(line, "")
		for _, rowData := range tokens[0 : len(tokens)-3] {
			rowMin, rowMax = findPlace(rowData, rowMin, rowMax)
		}
		for _, colData := range tokens[len(tokens)-3:] {
			colMin, colMax = findPlace(colData, colMin, colMax)
		}
		seatId := calculateId(rowMin, colMin)
		if seatId > maxSeatId {
			maxSeatId = seatId
		}
		seatMap[seatId] = true
	}

	t.Log(maxSeatId)

	for i := 0; i <= 127; i++ {
		for j := 0; j <= 7; j++ {
			previousSeat := seatMap[calculateId(i-1, j-1)]
			currentSeat := seatMap[calculateId(i, j)]
			nextSeat := seatMap[calculateId(i+1, j+1)]
			if previousSeat && nextSeat {
				if !currentSeat {
					t.Log(calculateId(i, j))
				}
			}
		}
	}
}

func calculateId(row, col int) int {
	return row*8 + col
}

func findPlace(place string, min, max int) (int, int) {
	if place == "F" || place == "L" {
		return min, (max + min) / 2
	}
	return ((max + min) / 2) + 1, max
}
