package aoc2020

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

/*
--- Day 2: Password Philosophy ---

Your flight departs in a few days from the coastal airport; the easiest way down to the coast from here is via toboggan.

The shopkeeper at the North Pole Toboggan Rental Shop is having a bad day. "Something's wrong with our computers; we can't log in!" You ask if you can take a look.

Their password database seems to be a little corrupted: some of the passwords wouldn't have been allowed by the Official Toboggan Corporate Policy that was in effect when they were chosen.

To try to debug the problem, they have created a list (your puzzle input) of passwords (according to the corrupted database) and the corporate policy when that password was set.

For example, suppose you have the following list:

1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc

Each line gives the password policy and then the password. The password policy indicates the lowest and highest number of times a given letter must appear for the password to be valid. For example, 1-3 a means that the password must contain a at least 1 time and at most 3 times.

In the above example, 2 passwords are valid. The middle password, cdefg, is not; it StringArrayContains no instances of b, but needs at least 1. The first and third passwords are valid: they contain one a or nine c, both within the limits of their respective policies.

How many passwords are valid according to their policies?
*/

type PasswordPolicy struct {
	minCount  int
	maxCount  int
	character string
}

func NewPasswordPolicy(min, max, character string) (*PasswordPolicy, error) {
	minI, err := strconv.Atoi(min)
	if err != nil {
		return nil, err
	}
	maxI, err := strconv.Atoi(max)
	if err != nil {
		return nil, err
	}

	return &PasswordPolicy{
		minCount:  minI,
		maxCount:  maxI,
		character: character,
	}, nil
}

type Password struct {
	data string
}

func (p *Password) validWithPolicy(policy *PasswordPolicy) bool {
	count := map[string]int{}
	for _, c := range p.data {
		count[string(c)]++
	}

	policyCharacterOccurence := count[policy.character]
	if policyCharacterOccurence >= policy.minCount && policyCharacterOccurence <= policy.maxCount {
		return true
	}

	return false
}

func (p *Password) validWithSecondPolicy(policy *PasswordPolicy) bool {
	password := strings.Split(p.data, "")
	pos1 := policy.minCount - 1
	pos2 := policy.maxCount - 1

	// XOR
	return (password[pos1] == policy.character || password[pos2] == policy.character) &&
		((password[pos1] == policy.character) != (password[pos2] == policy.character))
}

func parsePasswordPolicy(rx *regexp.Regexp, line string) (*Password, *PasswordPolicy) {
	submatch := rx.FindStringSubmatch(line)
	min := submatch[1]
	max := submatch[2]
	char := submatch[3]
	p, err := NewPasswordPolicy(min, max, char)
	if err != nil {
		log.Fatalln(err)
	}
	return &Password{data: submatch[4]}, p
}

func TestDay2(t *testing.T) {
	lines := ReadInputAsLines(2)
	parseRegex := regexp.MustCompile(`(\d+)-(\d+) (\w+): (\w+)`)

	validCount := 0
	validAlternatePolicyCount := 0
	for _, line := range lines {
		password, policy := parsePasswordPolicy(parseRegex, line)
		if password.validWithPolicy(policy) {
			validCount++
		}
		if password.validWithSecondPolicy(policy) {
			validAlternatePolicyCount++
		}
	}

	t.Log(validCount)
	t.Log(validAlternatePolicyCount)
}
