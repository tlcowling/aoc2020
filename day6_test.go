package aoc2020

import (
	"github.com/tlcowling/adventutils"
	"strings"
	"testing"
)

func TestDay6(t *testing.T) {
	lines := adventutils.ReadInputAsLines("./inputs/day6.txt")

	groupNumber := 0
	lineNumber := 0

	totals := make(map[int]map[string]int)
	peopleInGroup := make(map[int]int)

	for lineNumber < len(lines) {
		groupQuestionCount := make(map[string]int)
		personNumber := 0

		for lineNumber < len(lines) && lines[lineNumber] != "" {
			questionsAnswers := strings.Split(lines[lineNumber], "")
			for _, q := range questionsAnswers {
				groupQuestionCount[q]++
			}
			personNumber++
			lineNumber++
		}
		totals[groupNumber] = groupQuestionCount
		peopleInGroup[groupNumber] = personNumber
		groupNumber++
		lineNumber++
	}

	sum := 0
	for _, v := range totals {
		sum += len(v)
	}
	t.Log(sum)

	allAnsweredCountSum := 0
	for groupID, questionCount := range totals {
		allAnsweredCount := 0
		for _, answerCount := range questionCount {
			numPeopleInGroup := peopleInGroup[groupID]
			if answerCount == numPeopleInGroup {
				allAnsweredCount++
			}
		}
		allAnsweredCountSum += allAnsweredCount
	}
	t.Log(allAnsweredCountSum)
}
