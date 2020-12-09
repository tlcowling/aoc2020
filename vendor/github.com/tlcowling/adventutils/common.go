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
	return StringArrayToIntArray(ReadInputAsLines(path))
}

func ParseCommaSeparatedIntsFromFile(path string) []int {
	return ParseCommaSeparatedInts(fileContents(path))
}

func ParseCommaSeparatedInts(in string) []int {
	in = strings.TrimSpace(in)
	nums := strings.Split(in, ",")
	ints := make([]int, len(nums))
	for i, num := range nums {
		numI, err := strconv.Atoi(num)
		if err != nil {
			log.Fatalln(err)
		}
		ints[i] = numI
	}
	return ints
}

func StringArrayToIntArray(input []string) []int {
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

func ArrayMinMax(arr []int) (int, int) {
	min := arr[0]
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		el := arr[i]
		if el < min {
			min = el
		}
		if el >= max {
			max = el
		}
	}
	return min, max
}
