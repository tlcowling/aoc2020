package adventutils

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func ReadInputAsLines(path string) []string {
	return strings.Split(fileContents(path), "\n")
}

func fileContents(path string) string {
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	contents := string(fileBytes)
	contents = strings.TrimSpace(contents)
	return contents
}

func ReadInputAsInts(path string) []int {
	return inputStringsToInts(ReadInputAsLines(path))
}

func inputStringsToInts(input []string) []int {
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
