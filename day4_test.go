package aoc2020

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/tlcowling/adventutils"
	"strconv"
	"strings"
	"testing"
)

/*
You arrive at the airport only to realize that you grabbed your North Pole Credentials instead of your passport. While these documents are extremely similar, North Pole Credentials aren't issued by a country and therefore aren't actually valid documentation for travel in most of the world.

It seems like you're not the only one having problems, though; a very long line has formed for the automatic passport scanners, and the delay could upset your travel itinerary.

Due to some questionable network security, you realize you might be able to solve both of these problems at the same time.

The automatic passport scanners are slow because they're having trouble detecting which passports have all required fields. The expected fields are as follows:

    byr (Birth Year)
    iyr (Issue Year)
    eyr (Expiration Year)
    hgt (Height)
    hcl (Hair Color)
    ecl (Eye Color)
    pid (Passport ID)
    cid (Country ID)

Passport data is validated in batch files (your puzzle input). Each passport is represented as a sequence of key:value pairs separated by spaces or newlines. Passports are separated by blank lines.

Here is an example batch file containing four passports:

ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in

The first passport is valid - all eight fields are present. The second passport is invalid - it is missing hgt (the Height field).

The third passport is interesting; the only missing field is cid, so it looks like data from North Pole Credentials, not a passport at all! Surely, nobody would mind if you made the system temporarily ignore missing cid fields. Treat this "passport" as valid.

The fourth passport is missing two fields, cid and byr. Missing cid is fine, but missing any other field is not, so this passport is invalid.

According to the above rules, your improved system would report 2 valid passports.

Count the number of valid passports - those that have all required fields. Treat cid as optional. In your batch file, how many passports are valid?

*/
func TestDay4(t *testing.T) {
	lines := adventutils.ReadInputAsLines("./inputs/day4.txt")

	validPassportCount := 0
	validPassportValuesCount := 0
	for i := 0; i < len(lines); i++ {
		passport := make(map[string]string)
		for i < len(lines) {
			line := lines[i]
			if line == "" {
				break
			}
			pairs := strings.Split(line, " ")
			for _, pair := range pairs {
				tokens := strings.Split(pair, ":")
				passport[tokens[0]] = tokens[1]
			}

			i++
		}

		if hasExpectedKeys(passport) {
			validPassportCount++
			if hasValidFieldValues(passport) {
				validPassportValuesCount++
			}
		}

	}
	t.Log(validPassportCount)
	t.Log(validPassportValuesCount)
}

func hasExpectedKeys(passport map[string]string) bool {
	expected := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
		//"cid", //(optional)
	}

	for _, e := range expected {
		if passport[e] == "" {
			return false
		}
	}

	return true
}

func validateFieldIntValues(field string, min, max int) error {
	if field == "" {
		return errors.New("err: field is empty")
	}
	atoi, err := strconv.Atoi(field)
	if err != nil {
		return fmt.Errorf("err: field %v is invalid", field)
	}
	if !(atoi >= min && atoi <= max) {
		return fmt.Errorf("err: field %v is outside range [%d-%d]", field, min, max)
	}
	return nil
}

func hasValidFieldValues(passport map[string]string) bool {
	if validateFieldIntValues(passport["byr"], 1920, 2002) != nil {
		return false
	}
	if validateFieldIntValues(passport["iyr"], 2010, 2020) != nil {
		return false
	}
	if validateFieldIntValues(passport["eyr"], 2020, 2030) != nil {
		return false
	}

	height := passport["hgt"]
	if height == "" {
		fmt.Println("invalid height")
		return false
	}
	number, units := height[0:len(height)-2], height[len(height)-2:len(height)]
	switch units {
	case "cm":
		if validateFieldIntValues(number, 150, 193) != nil {
			return false
		}
	case "in":
		if validateFieldIntValues(number, 59, 76) != nil {
			return false
		}
	default:
		return false
	}

	haircolor := passport["hcl"]
	if haircolor == "" || !strings.HasPrefix(haircolor, "#") {
		return false
	}
	haircolor = haircolor[1:]
	_, err := hex.DecodeString(haircolor)
	if err != nil {
		fmt.Println("invalid hcl")
		return false
	}

	allowedEyeColors := []string{
		"amb", "blu", "brn", "gry", "grn", "hzl", "oth",
	}
	eyeColor := passport["ecl"]

	if !adventutils.StringArrayContains(allowedEyeColors, eyeColor) {
		//fmt.Println("invalid ecl", eyeColor)
		return false
	}

	passid := passport["pid"]
	if len(passid) != 9 {
		return false
	}

	//after ALL THAT
	return true
}
