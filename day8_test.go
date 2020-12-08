package aoc2020

import (
	"github.com/stretchr/testify/assert"
	"github.com/tlcowling/adventutils"

	"log"
	"strconv"
	"strings"
	"testing"
)

const (
	noop       = "nop"
	jump       = "jmp"
	accumulate = "acc"
)

func TestDay8(t *testing.T) {
	lines := adventutils.ReadInputAsLines("./inputs/day8.txt")
	_, acc, _ := RunInstruction(lines)
	t.Log(acc)

	_, newacc, _ := RunModifiedInstructions(lines, len(lines))
	t.Log(newacc)
}

func RunInstruction(program []string) (int, int, bool) {
	return runInstructions(program, -1)
}

func RunModifiedInstructions(program []string, modificationCount int) (int, int, bool) {
	for i := 0; i < modificationCount; i++ {
		lastInstruction, acc, infinite := runInstructions(program, i)
		if !infinite {
			return lastInstruction, acc, infinite
		}
	}
	return 0, 0, false
}

func switchInstruction(in string) string {
	switch in {
	case jump:
		return noop
	case noop:
		return jump
	}
	return in
}

func runInstructions(lines []string, instructionPointerModifyLine int) (int, int, bool) {
	visits := make(map[int]int)
	acc := 0
	instructionPointer := 0

	for instructionPointer < len(lines) {
		visits[instructionPointer]++
		if visits[instructionPointer] > 1 {
			return instructionPointer, acc, true
		}
		instruction, amount := parseInstruction(lines[instructionPointer])
		if instructionPointerModifyLine == instructionPointer {
			instruction = switchInstruction(instruction)
		}
		switch instruction {
		case noop:
		case jump:
			instructionPointer += amount
			continue
		case accumulate:
			acc += amount
		}
		instructionPointer++
	}
	return instructionPointer, acc, false
}

func parseInstruction(in string) (string, int) {
	split := strings.Split(in, " ")
	i, err := strconv.Atoi(split[1])
	if err != nil {
		log.Fatalln(err)
	}
	return split[0], i
}

func TestParseLine(t *testing.T) {
	inst, i := parseInstruction("jmp +4")
	assert.Equal(t, inst, "jmp")
	assert.Equal(t, i, 4)

	inst, i = parseInstruction("acc -204")
	assert.Equal(t, inst, "acc")
	assert.Equal(t, i, -204)
}
