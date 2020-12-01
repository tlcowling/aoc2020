package aoc2020

import (
	"fmt"
	"io/ioutil"
	"log"
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
