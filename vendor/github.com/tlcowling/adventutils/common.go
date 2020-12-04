package adventutils

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func ReadInputAsLines(path string) []string {
	all, err := ioutil.ReadFile(path)
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

func StringArrayContains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
