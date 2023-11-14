package day01

import "joenace.com/aoc-2015-go/helpers"

var input string

func init() {
	input = helpers.ReadInput()
}

func Part1() int {
	return solve1(input)
}

func Part2() int {
	return solve2(input)
}

func solve1(input string) int {
	sum := 0

	for _, value := range input {
		if value == '(' {
			sum++
		} else if value == ')' {
			sum--
		}
	}

	return sum
}

func solve2(input string) int {
	sum := 0

	for i, value := range input {
		if value == '(' {
			sum++
		} else if value == ')' {
			sum--

			if sum < 0 {
				return i + 1
			}
		}
	}

	return 0
}
