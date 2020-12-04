package aoc2020

import (
	"github.com/tlcowling/adventutils"
	"testing"
)

/*
--- Day 1: Report Repair ---

After saving Christmas five years in a row, you've decided to take a vacation at a nice resort on a tropical island. Surely, Christmas will go on without you.

The tropical island has its own currency and is entirely cash-only. The gold coins used there have a little picture of a starfish; the locals just call them stars. None of the currency exchanges seem to have heard of them, but somehow, you'll need to find fifty of these coins by the time you arrive so you can pay the deposit on your room.

To save your vacation, you need to get all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

Before you leave, the Elves in accounting just need you to fix your expense report (your puzzle input); apparently, something isn't quite adding up.

Specifically, they need you to find the two entries that sum to 2020 and then multiply those two numbers together.

For example, suppose your expense report contained the following:

1721
979
366
299
675
1456

In this list, the two entries that sum to 2020 are 1721 and 299. Multiplying them together produces 1721 * 299 = 514579, so the correct answer is 514579.

Of course, your expense report is much larger. Find the two entries that sum to 2020; what do you get if you multiply them together?
/* ORIGINAL
func TestDay1(t *testing.T) {
	input := ReadInputAsLines(1)
	inputInts := InputStringsToInts(input)

	for i := 0; i < len(inputInts)-1; i++ {
		for j := i + 1; j < len(inputInts); j++ {
			if inputInts[i]+inputInts[j] == 2020 {
				t.Log("Day 1 Part 1")
				t.Log("Numbers:", inputInts[i], inputInts[j])
				t.Log("Multiplied:", inputInts[i]*inputInts[j])
			}
		}
	}
}

--- Part Two ---

The Elves in accounting are thankful for your help; one of them even offers you a starfish coin they had left over from a past vacation. They offer you a second one if you can find three numbers in your expense report that meet the same criteria.

Using the above example again, the three entries that sum to 2020 are 979, 366, and 675. Multiplying them together produces the answer, 241861950.

In your expense report, what is the product of the three entries that sum to 2020?

func TestDayOneP2(t *testing.T) {
	input := ReadInputAsLines(1)
	inputInts := InputStringsToInts(input)

	for i := 0; i < len(inputInts)-2; i++ {
		for j := i + 1; j < len(inputInts)-1; j++ {
			for k := j + 2; k < len(inputInts); k++ {
				first := inputInts[i]
				next := inputInts[j]
				another := inputInts[k]

				if first+next+another == 2020 {
					t.Log("Day 1 Part 2")
					t.Log("Numbers:", first, next, another)
					t.Log("Multiplied:", first*next*another)
				}
			}
		}
	}
}
*/

/*
Add optimisation for only needing to loop twice but with extra storage
- store precalculated sums and products in hashmap (since after summing we would lose information about the product)
- no point storing first number if sum greater than 2020
*/
func TestDay1(t *testing.T) {
	inputInts := adventutils.ReadInputAsInts("./inputs/day1.txt")
	preCalculated := make(map[int]int)

	p1Product := 0
	for i := 0; i < len(inputInts)-2; i++ {
		for j := i + 1; j < len(inputInts)-1; j++ {
			if inputInts[i]+inputInts[j] == 2020 {
				p1Product = inputInts[i] * inputInts[j]
			}
			if inputInts[i]+inputInts[j] <= 2020 {
				preCalculated[inputInts[i]+inputInts[j]] = inputInts[i] * inputInts[j]
			}
		}
	}

	t.Log(p1Product)

	for preCalculatedSum, preCalculatedProduct := range preCalculated {
		for _, inputInt := range inputInts {
			if preCalculatedSum+inputInt == 2020 {
				p2Product := preCalculatedProduct * inputInt
				t.Log(p2Product)
				return
			}
		}
	}
}
