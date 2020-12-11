package aoc2020

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tlcowling/adventutils"
	"reflect"
	"strings"
	"testing"
)

const (
	SeatEmpty    = "L"
	SeatOccupied = "#"
	NoSeat       = "."
)

func TestDay11(t *testing.T) {
	ferryRows := adventutils.ReadInputAsLines("./inputs/day11.txt")
	seats := makeSeatGrid(ferryRows)
	//printGrid(seats)

	previousPhase := seats
	var steadyState bool
	for !steadyState {
		np := nextPhase(previousPhase)
		//printGrid(np)
		if reflect.DeepEqual(np, previousPhase) {
			steadyState = true
		}
		previousPhase = np
	}
	t.Log(countsSeatsInState(allSeats(previousPhase), SeatOccupied))

	previousPhase = seats
	steadyState = false
	for !steadyState {
		np := nextPhaseDay2(previousPhase)
		//printGrid(np)
		if reflect.DeepEqual(np, previousPhase) {
			steadyState = true
		}
		previousPhase = np

	}
	printGrid(previousPhase)
	t.Log(countsSeatsInState(allSeats(previousPhase), SeatOccupied))
}
func allSeats(g [][]string) []string {
	allSeats := []string{}
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			allSeats = append(allSeats, g[i][j])
		}
	}
	return allSeats
}

func nextPhase(seats [][]string) [][]string {
	next := makeGridCopy(seats)
	for rowID := 0; rowID < len(seats); rowID++ {
		for columnID := 0; columnID < len(seats[0]); columnID++ {
			targetSeat := seats[rowID][columnID]
			adjacentSeats := adjacentSeats(seats, rowID, columnID)
			switch targetSeat {
			case SeatEmpty:
				occupiedAdjacentSeats := countsSeatsInState(adjacentSeats, SeatOccupied)
				if occupiedAdjacentSeats == 0 {
					next[rowID][columnID] = "#"
				}
			case SeatOccupied:
				occupiedAdjacentSeats := countsSeatsInState(adjacentSeats, SeatOccupied)
				if occupiedAdjacentSeats >= 4 {
					next[rowID][columnID] = "L"
				}
			case NoSeat:
				// nochange
			}
		}
	}

	return next
}

func nextPhaseDay2(seats [][]string) [][]string {
	next := makeGridCopy(seats)
	for rowID := 0; rowID < len(seats); rowID++ {
		for columnID := 0; columnID < len(seats[0]); columnID++ {
			targetSeat := seats[rowID][columnID]
			visibleSeats := visibleSeats(seats, rowID, columnID)
			switch targetSeat {
			case SeatEmpty:
				occupiedVisibleSeats := countsSeatsInState(visibleSeats, SeatOccupied)
				if occupiedVisibleSeats == 0 {
					next[rowID][columnID] = "#"
				}
			case SeatOccupied:
				occupiedVisibleSeats := countsSeatsInState(visibleSeats, SeatOccupied)
				if occupiedVisibleSeats >= 5 {
					next[rowID][columnID] = "L"
				}
			case NoSeat:
				// nochange
			}
		}
	}

	return next
}

func printGrid(g [][]string) {
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[0]); j++ {
			fmt.Printf(g[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func TestGridCopy(t *testing.T) {
	original := [][]string{
		[]string{"1", "2", "3"},
		[]string{"4", "5", "6"},
		[]string{"7", "8", "9"},
	}
	newGrid := makeGridCopy(original)
	assert.Equal(t, newGrid, original)
	original[0] = []string{"10", "11", "12"}
	assert.NotEqual(t, newGrid, original)
}

func makeGridCopy(grid [][]string) [][]string {
	n := len(grid)
	g := make([][]string, n)

	for i := 0; i < n; i++ {
		newRow := []string{}
		for j := 0; j < len(grid[i]); j++ {
			newRow = append(newRow, grid[i][j])
		}

		g[i] = newRow
	}
	return g
}

func countsSeatsInState(seats []string, state string) int {
	seatCount := 0
	for _, s := range seats {
		if s == state {
			seatCount++
		}
	}
	return seatCount
}

func TestVisibleSeats(t *testing.T) {
	grid1 := [][]string{
		strings.Split(".............", ""),
		strings.Split(".L.L.#.#.#.#.", ""),
		strings.Split(".............", ""),
	}
	assert.Equal(t, countsSeatsInState(visibleSeats(grid1, 1, 1), SeatEmpty), 1)
	assert.Equal(t, countsSeatsInState(visibleSeats(grid1, 1, 1), SeatOccupied), 0)

	grid2 := [][]string{
		strings.Split(".##.##.", ""),
		strings.Split("#.#.#.#", ""),
		strings.Split("##...##", ""),
		strings.Split("...L...", ""),
		strings.Split("##...##", ""),
		strings.Split("#.#.#.#", ""),
		strings.Split(".##.##.", ""),
	}
	assert.Equal(t, countsSeatsInState(visibleSeats(grid2, 3, 3), SeatEmpty), 0)
}

func visibleSeats(grid [][]string, rowID, columnID int) []string {
	seats := []string{}

	trackRow := rowID
	trackCol := columnID
	//north
	if trackRow > 0 {
		//fmt.Println(trackRow, trackCol)
		trackRow--
		seenSeat := grid[trackRow][columnID]
		for keepSearchingForSeat(grid, trackCol, trackRow, seenSeat) {
			//fmt.Println(trackRow, trackCol)
			seenSeat = grid[trackRow][trackCol]
			trackRow--
		}
		seats = append(seats, seenSeat)
	}
	// south
	trackRow = rowID
	trackCol = columnID
	if trackRow < len(grid)-1 {
		trackRow++
		seenSeat := grid[trackRow][columnID]
		for keepSearchingForSeat(grid, trackCol, trackRow, seenSeat) {
			//fmt.Println(trackRow, trackCol)

			seenSeat = grid[trackRow][trackCol]
			trackRow++
		}
		seats = append(seats, seenSeat)
	}
	//fmt.Println("east")
	// east
	trackRow = rowID
	trackCol = columnID
	if trackCol < len(grid)-1 {
		trackCol++
		seenSeat := grid[trackRow][trackCol]
		for keepSearchingForSeat(grid, trackCol, trackRow, seenSeat) {
			//fmt.Println(trackRow, trackCol)

			seenSeat = grid[trackRow][trackCol]
			trackCol++
		}
		seats = append(seats, seenSeat)
	}

	//fmt.Println("west")
	trackRow = rowID
	trackCol = columnID
	if trackCol > 0 {
		trackCol--
		seenSeat := grid[trackRow][trackCol]

		for keepSearchingForSeat(grid, trackCol, trackRow, seenSeat) {
			//fmt.Println("west", trackRow, trackCol)
			seenSeat = grid[trackRow][trackCol]
			trackCol--
		}
		seats = append(seats, seenSeat)
	}

	//fmt.Println("south east")
	trackRow = rowID
	trackCol = columnID
	if trackCol < len(grid[0])-1 && trackRow < len(grid)-1 {
		trackCol++
		trackRow++
		//fmt.Println("trouble at", trackRow, trackCol)
		seenSeat := grid[trackRow][trackCol]

		for keepSearchingForSeat(grid, trackCol, trackRow, seenSeat) {
			//fmt.Println("south east", trackRow, trackCol)
			seenSeat = grid[trackRow][trackCol]
			trackCol++
			trackRow++
		}
		seats = append(seats, seenSeat)
	}

	//fmt.Println("south west")
	trackRow = rowID
	trackCol = columnID
	if trackCol > 0 && trackRow < len(grid)-1 {
		trackCol--
		trackRow++
		seenSeat := grid[trackRow][trackCol]

		for keepSearchingForSeat(grid, trackCol, trackRow, seenSeat) {
			//fmt.Println("south west", trackRow, trackCol)
			seenSeat = grid[trackRow][trackCol]
			trackCol--
			trackRow++
		}
		seats = append(seats, seenSeat)
	}

	//fmt.Println("north west")
	trackRow = rowID
	trackCol = columnID
	if trackRow > 0 && columnID > 0 {
		trackCol--
		trackRow--
		seenSeat := grid[trackRow][trackCol]

		for keepSearchingForSeat(grid, trackCol, trackRow, seenSeat) {
			//fmt.Println("north west", trackRow, trackCol)
			seenSeat = grid[trackRow][trackCol]
			trackCol--
			trackRow--
		}
		seats = append(seats, seenSeat)
	}

	//fmt.Println("north east")
	trackRow = rowID
	trackCol = columnID
	if trackRow > 0 && columnID < len(grid[0])-1 {
		trackCol++
		trackRow--
		seenSeat := grid[trackRow][trackCol]

		for keepSearchingForSeat(grid, trackCol, trackRow, seenSeat) {
			//fmt.Println("north east", trackRow, trackCol)
			seenSeat = grid[trackRow][trackCol]
			trackCol++
			trackRow--
		}
		seats = append(seats, seenSeat)
	}

	return seats
}

func keepSearchingForSeat(grid [][]string, trackCol int, trackRow int, seenSeat string) bool {
	return trackCol >= 0 && trackCol < len(grid[0]) &&
		trackRow >= 0 && trackRow < len(grid) && seenSeat == "."
}

func adjacentSeats(grid [][]string, rowID, columnID int) []string {
	// . . .
	// . X .
	// . . .
	columnPositions := []int{
		columnID - 1, columnID, columnID + 1,
		columnID - 1, columnID + 1,
		columnID - 1, columnID, columnID + 1,
	}
	rowPositions := []int{
		rowID - 1, rowID - 1, rowID - 1,
		rowID, rowID,
		rowID + 1, rowID + 1, rowID + 1,
	}
	adjacents := []string{}
	for i := 0; i < len(columnPositions); i++ {
		// position outside grid
		if rowPositions[i] < 0 || rowPositions[i] >= len(grid) || columnPositions[i] < 0 || columnPositions[i] >= len(grid[0]) {
			continue
		}

		adjacents = append(adjacents, grid[rowPositions[i]][columnPositions[i]])
	}
	return adjacents
}

func makeSeatGrid(lines []string) [][]string {
	seatGrid := make([][]string, len(lines))

	for rowID, row := range lines {
		rowGrid := make([]string, len(row))
		seats := strings.Split(row, "")
		for columnID, seat := range seats {
			rowGrid[columnID] = seat
		}
		seatGrid[rowID] = rowGrid
	}
	return seatGrid
}
