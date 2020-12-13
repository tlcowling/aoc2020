package aoc2020

import (
	"github.com/stretchr/testify/assert"
	"github.com/tlcowling/adventutils"
	"log"
	"strconv"
	"testing"
)

const (
	DirectionNorth = iota
	DirectionEast
	DirectionSouth
	DirectionWest
)

type Coordinate struct {
	x int
	y int
}

func (c *Coordinate) ManhattanFromCoordinate(c2 *Coordinate) int {
	return adventutils.AbsInt(c.y-c2.y) + adventutils.AbsInt(c.x-c2.x)
}

func TestManhattan(t *testing.T) {
	origin := &Coordinate{0, 0}
	other := &Coordinate{10, -10}
	other2 := &Coordinate{-10, -10}
	assert.Equal(t, 0, origin.ManhattanFromCoordinate(&Coordinate{0, 0}), "distance from self ")
	assert.Equal(t, 20, other.ManhattanFromCoordinate(&Coordinate{0, 0}), "distance from other to self ")
	assert.Equal(t, 20, other2.ManhattanFromCoordinate(other), "distance from other to other2")
	assert.Equal(t, 20, other.ManhattanFromCoordinate(other2), "reverse ok")
}

type Ferry struct {
	pos       *Coordinate
	direction int
	waypoint  *Coordinate
}

func NewFerry(direction int) *Ferry {
	return &Ferry{
		pos:       &Coordinate{0, 0},
		direction: direction,
	}
}

func (f *Ferry) parseInstruction(in string) (string, int) {
	dir := in[0:1]
	rest := in[1:]
	quantity, err := strconv.Atoi(rest)
	if err != nil {
		log.Fatalln(err)
	}
	return dir, quantity
}

func (f *Ferry) moveShip(in string) {
	dir, quantity := f.parseInstruction(in)

	switch dir {
	case "R":
		f.rotate(quantity)
	case "L":
		f.rotate(-quantity)
	case "F":
		f.move(quantity)
	case "N":
		f.pos.y += quantity
	case "S":
		f.pos.y -= quantity
	case "E":
		f.pos.x += quantity
	case "W":
		f.pos.x -= quantity
	}
}

func (f *Ferry) moveWithWaypoint(in string) {
	dir, quantity := f.parseInstruction(in)

	switch dir {
	case "R":
		f.rotate(quantity)
	case "L":
		f.rotate(-quantity)
	case "F":
		f.move(quantity)
	case "N":
		f.waypoint.y += quantity
	case "S":
		f.waypoint.y -= quantity
	case "E":
		f.waypoint.x += quantity
	case "W":
		f.waypoint.x -= quantity
	}
}

func (f *Ferry) rotateWayPoint(deg int) {
	//turns := deg/90
	//rotations:=modImUsedTo(turns, 4)
	//switch rotations {
	//case 0:
	//	//nothing
	//case 1:
	//	tmpX := f.waypoint.x
	//	f.waypoint.x = f.waypoint.y
	//}

}

func (f *Ferry) move(q int) {
	switch f.direction {
	case DirectionNorth:
		f.pos.y += q
	case DirectionSouth:
		f.pos.y -= q
	case DirectionEast:
		f.pos.x += q
	case DirectionWest:
		f.pos.x -= q
	}
}

func (f *Ferry) rotate(degs int) {
	turns := degs / 90

	f.direction += turns
	f.direction = modImUsedTo(f.direction, 4)
}

func modImUsedTo(div, mod int) int {
	res := div % mod
	if (res < 0 && mod > 0) || (res > 0 && mod < 0) {
		return res + mod
	}
	return res
}

func TestModImUsedTo(t *testing.T) {
	assert.Equal(t, 0, modImUsedTo(0, 4))
	assert.Equal(t, 3, modImUsedTo(-1, 4))
	assert.Equal(t, 1, modImUsedTo(5, 4))
	assert.Equal(t, 0, modImUsedTo(4, 4))
}

func TestDay12(t *testing.T) {
	f := NewFerry(DirectionEast)
	lines := adventutils.ReadInputAsLines("./inputs/day12_test.txt")
	for _, instruction := range lines {
		f.moveShip(instruction)
		//fmt.Println(i, f.direction, f.pos)
	}
	t.Log(f.pos.ManhattanFromCoordinate(&Coordinate{0, 0}))

	//f2 := NewFerry(DirectionEast)
	//f2.waypoint = &Coordinate{10, 1}
	//for _, instruction := range lines {
	//	f.moveWithWaypoint(instruction)
	//	//fmt.Println(i, f.direction, f.pos)
	//}
	//t.Log(f.pos.ManhattanFromCoordinate(&Coordinate{0,0}) )
}
