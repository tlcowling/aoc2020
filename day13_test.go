package aoc2020

import (
	"github.com/tlcowling/adventutils"
	"log"
	"math"
	"strconv"
	"strings"
	"testing"
)

func TestDay13(t *testing.T) {
	lines := adventutils.ReadInputAsLines("./inputs/day13.txt")
	timestampStr := lines[0]
	timestamp, err := strconv.Atoi(timestampStr)
	if err != nil {
		log.Fatalln(err)
	}
	buses := lines[1]

	allBuses := make(map[int]int)
	busTimestamps := strings.Split(buses, ",")
	for _, b := range busTimestamps {
		if b == "x" {
			continue
		}
		bus, err := strconv.Atoi(b)
		if err != nil {
			log.Fatalln(err)
		}

		allBuses[bus] = eachBus(bus, timestamp)
	}

	closestBus := timestamp * 2
	var closestBusID int
	for bID, b := range allBuses {
		busWait := b - timestamp
		if busWait < closestBus {
			closestBus = busWait
			closestBusID = bID
		}
	}
	t.Log(closestBus * closestBusID)
}

func eachBus(id, timestamp int) int {
	multiple := float64(timestamp) / float64(id)
	round := math.Ceil(multiple)
	closest := int(round) * id
	return closest
}
