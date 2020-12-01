package aoc2020

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func ReadInputAsLines(day int) []string {
	all, err := ioutil.ReadFile(fmt.Sprintf("./inputs/day%d.txt", day))
	if err != nil {
		log.Fatalln(err)
	}
	contents := string(all)
	contents = strings.TrimSpace(contents)
	return strings.Split(contents, "\n")
}

func InputStringsToInts(input []string) []int {
	ints := make([]int, len(input))
	for i, str := range input {
		atoi, err := strconv.Atoi(str)
		if err != nil {
			log.Fatalln(err)
		}
		ints[i] = atoi
	}
	return ints
}
